package gohttp

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/printer"
)

var (
	firewallPassing bool
)

// HTTPOptions :: This is Struct Http, it will inherit the struct Options and client.
type httpOptions struct {
	url                  *URLOptions
	method               string
	tlsCertificateVerify bool
	tor                  bool
	proxy                func(*http.Request) (*url.URL, error)
	data                 io.Reader
	userAgent            string
	redirect             func(req *http.Request, via []*http.Request) error
	contentType          string
	firewall             bool
	sleep                time.Duration
}

// Response :: This struct will store the request data, and will be used for a return.
type Response struct {
	RawIo    io.Reader
	Raw      string
	URL      *URLOptions
	Response *http.Response
}

// URLOptions :: This struct will be used to inform directories ... the complete URL ... or just the domain.
// (Alert) The focus of this struct is to be used together with HTTPOptions!
type URLOptions struct {
	Simple    string
	Full      string
	Directory string
	URL       *url.URL
}

// SimpleRequest :: The first parameter must always be the base url. The second must be the directory.
func SimpleRequest(params ...string) *Response {
	http := NewHTTPClient()
	http.SetURL(params[0])

	if len(params) > 1 {
		http.SetURLDirectory(params[1])
	}

	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Fatal(err)
	}

	return response
}

func NewHTTPClient() *httpOptions {
	options := &httpOptions{
		method:      "GET",
		userAgent:   "WPrecon - Wordpress Recon (Vulnerability Scanner)",
		data:        nil,
		contentType: "text/html; charset=UTF-8"}

	options.url = &URLOptions{}

	return options
}

func (options *httpOptions) SetURL(url string) *httpOptions {
	if !strings.HasSuffix(url, "/") {
		options.url.Simple = url + "/"
		options.url.Full = url + "/"
	} else {
		options.url.Simple = url
		options.url.Full = url
	}

	return options
}

func (options *httpOptions) SetURLDirectory(directory string) *httpOptions {
	if !strings.HasPrefix(directory, "/") && !strings.HasSuffix(options.url.Simple, "/") {
		options.url.Directory = "/" + directory
		options.url.Full = options.url.Simple + "/" + directory
	} else {
		options.url.Directory = directory
		options.url.Full = options.url.Simple + directory
	}

	return options
}

func (options *httpOptions) SetURLFull(full string) *httpOptions {
	options.url.Full = full

	return options
}

func (options *httpOptions) OnTor(status bool) (*httpOptions, error) {
	if status {
		tor, err := url.Parse("http://127.0.0.1:9080")

		if err != nil {
			return nil, fmt.Errorf("proxy URL is invalid (%w)", err)
		}

		options.proxy = http.ProxyURL(tor)
	}

	return options, nil
}

func (options *httpOptions) OnRandomUserAgent(status bool) *httpOptions {
	if status {
		options.userAgent = randomuseragent()
	}

	return options
}

func (options *httpOptions) OnTLSCertificateVerify(status bool) *httpOptions {
	options.tlsCertificateVerify = status

	return options
}

func (options *httpOptions) SetMethod(method string) *httpOptions {
	options.method = method

	return options
}

func (options *httpOptions) SetUserAgent(userAgent string) *httpOptions {
	options.userAgent = userAgent

	return options
}

func (options *httpOptions) SetForm(form *url.Values) *httpOptions {
	options.data = strings.NewReader(form.Encode())

	return options
}

func (options *httpOptions) SetData(data string) *httpOptions {
	options.data = strings.NewReader(data)

	return options
}

func (options *httpOptions) SetRedirectFunc(redirectFunc func(req *http.Request, via []*http.Request) error) *httpOptions {
	options.redirect = redirectFunc

	return options
}

func (options *httpOptions) SetContentType(contentType string) *httpOptions {
	options.contentType = contentType

	return options
}

func (options *httpOptions) FirewallDetection(status bool) *httpOptions {
	options.firewall = status

	return options
}

func (options *httpOptions) SetSleep(tm int) *httpOptions {
	options.sleep = time.Duration(tm)

	return options
}

func (options *httpOptions) f(http *Response) {
	exists, firewall, output, solve, confidence := NewFirewallDetectionPassive(http).All().Run()

	if exists {
		printer.Danger("Firewall Active Detection:")
		printer.List("Firewall:", firewall).D()
		printer.List("Detection By:", output).D()
		printer.List("Confidence:", fmt.Sprintf("%d%%", confidence)).D()
		if solve != "" {
			printer.List("Solve:", solve).Warning()
		}

		if response := printer.ScanQ("Do you wish to continue ? [y]es | [N]o : "); response != "y" {
			printer.Fatal("Exiting...")
		}

		printer.Println()
		firewallPassing = true
	}
}

func (options *httpOptions) Run() (*Response, error) {
	client := &http.Client{
		CheckRedirect: options.redirect,
		Transport: &http.Transport{
			Proxy:             options.proxy,
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: options.tlsCertificateVerify,
			},
		},
	}

	request, err := http.NewRequest(options.method, options.url.Full, options.data)

	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", options.userAgent)
	request.Header.Set("Content-Type", options.contentType)

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	Database.TotalRequests++

	structResponse := &Response{
		Raw:      string(raw),
		URL:      options.url,
		Response: response,
	}

	if options.firewall && !firewallPassing {
		options.f(structResponse)
	}

	if options.sleep != 0 {
		time.Sleep(time.Duration(options.sleep) * time.Second)
	} else if sleep := Database.OtherInformationsInt["http.requests.time.sleep"]; sleep != 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	return structResponse, nil
}
