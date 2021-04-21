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

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

var (
	firewallPass bool
)

// httpOptions :: This is Struct Http, it will inherit the struct self and client.
type httpOptions struct {
	url                  *URLOptions
	method               string
	tlsCertificateVerify bool
	// tor                  bool
	proxy       func(*http.Request) (*url.URL, error)
	data        io.Reader
	userAgent   string
	redirect    func(req *http.Request, via []*http.Request) error
	contentType string
	firewall    bool
	sleep       time.Duration
}

// Response :: This struct will store the request data, and will be used for a return.
type Response struct {
	RawIo    io.Reader
	Raw      string
	URL      *URLOptions
	Response *http.Response
}

// URLOptions :: This struct will be used to inform directories... the complete URL... or just the domain.
// (Alert) The focus of this struct is to be used together with httpOptions!
type URLOptions struct {
	Simple    string
	Full      string
	Directory string
	URL       *url.URL
}

// SimpleRequest :: The first parameter must always be the base url. The second must be the directory.
func SimpleRequest(params ...string) *Response {
	var http = NewHTTPClient()

	http.SetURL(params[0])

	if len(params) > 1 {
		http.SetURLDirectory(strings.Join(params[1:], ""))
	}

	http.OnTor(database.Memory.GetBool("HTTP self TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP self Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP self TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Fatal(err)
	}

	return response
}

func SimpleRequestGoroutine(channel chan *Response, params ...string) {
	var http = NewHTTPClient()

	http.SetURL(params[0])

	if len(params) > 1 {
		http.SetURLDirectory(strings.Join(params[1:], ""))
	}

	http.OnTor(database.Memory.GetBool("HTTP self TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP self Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP self TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Fatal(err)
	}

	channel <- response
}

func NewHTTPClient() *httpOptions {
	self := &httpOptions{
		method:      "GET",
		userAgent:   "WPrecon - Wordpress Recon (Vulnerability Scanner)",
		data:        nil,
		contentType: "text/html; charset=UTF-8"}

	self.url = &URLOptions{}

	return self
}

func (self *httpOptions) SetURL(url string) *httpOptions {
	if !strings.HasSuffix(url, "/") {
		self.url.Simple = url + "/"
		self.url.Full = url + "/"
	} else {
		self.url.Simple = url
		self.url.Full = url
	}

	return self
}

func (self *httpOptions) SetURLDirectory(directory string) *httpOptions {
	if !strings.HasPrefix(directory, "/") && !strings.HasSuffix(self.url.Simple, "/") {
		self.url.Directory = "/" + directory
		self.url.Full = self.url.Simple + "/" + directory
	} else {
		self.url.Directory = directory
		self.url.Full = self.url.Simple + directory
	}

	return self
}

func (self *httpOptions) SetURLFull(full string) *httpOptions {
	self.url.Full = full

	return self
}

func (self *httpOptions) OnTor(status bool) (*httpOptions, error) {
	if status {
		tor, err := url.Parse("http://127.0.0.1:9080")

		if err != nil {
			return nil, fmt.Errorf("proxy URL is invalid (%w)", err)
		}

		self.proxy = http.ProxyURL(tor)
	}

	return self, nil
}

func (self *httpOptions) OnRandomUserAgent(status bool) *httpOptions {
	if status {
		self.userAgent = RandomUserAgent()
	}

	return self
}

func (self *httpOptions) OnTLSCertificateVerify(status bool) *httpOptions {
	self.tlsCertificateVerify = status

	return self
}

func (self *httpOptions) SetMethod(method string) *httpOptions {
	self.method = method

	return self
}

func (self *httpOptions) SetUserAgent(userAgent string) *httpOptions {
	self.userAgent = userAgent

	return self
}

func (self *httpOptions) SetForm(form *url.Values) *httpOptions {
	self.data = strings.NewReader(form.Encode())

	return self
}

func (self *httpOptions) SetData(data string) *httpOptions {
	self.data = strings.NewReader(data)

	return self
}

func (self *httpOptions) SetRedirectFunc(redirectFunc func(req *http.Request, via []*http.Request) error) *httpOptions {
	self.redirect = redirectFunc

	return self
}

func (self *httpOptions) SetContentType(contentType string) *httpOptions {
	self.contentType = contentType

	return self
}

func (self *httpOptions) FirewallDetection(status bool) *httpOptions {
	self.firewall = status

	return self
}

func (self *httpOptions) SetSleep(tm int) *httpOptions {
	self.sleep = time.Duration(tm)

	return self
}

func (self *httpOptions) FirewallActiveDetection(http *Response) {
	exists, firewall, output, solve, confidence := NewFirewallDetectionPassive(http).All().Run()

	if exists {
		printer.Danger("Firewall Active Detection:", firewall)
		printer.NewTopics("Detection By:", output).Default()
		printer.NewTopics("Confidence:", fmt.Sprintf("%d%%", confidence)).Default()
		if solve != "" {
			printer.NewTopics("Solve:", solve).Warning()
		}

		if response := printer.ScanQ("Do you wish to continue ? [y]es | [N]o : "); response != "y" && response != "\n" {
			printer.Fatal("Exiting...")
		}

		printer.Println()
		firewallPass = true
	}
}

func (self *httpOptions) Run() (*Response, error) {
	client := &http.Client{
		CheckRedirect: self.redirect,
		Transport: &http.Transport{
			Proxy: self.proxy,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: self.tlsCertificateVerify,
			},
		},
	}

	request, err := http.NewRequest(self.method, self.url.Full, self.data)

	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", self.userAgent)
	request.Header.Set("Content-Type", self.contentType)

	response, err := client.Do(request)

	if err != nil {
		defer response.Body.Close()

		return nil, err
	}

	defer response.Body.Close()

	raw, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	database.Memory.AddInt("HTTP Total")

	structResponse := &Response{
		Raw:      string(raw),
		URL:      self.url,
		Response: response,
	}

	if self.firewall && !firewallPass {
		self.FirewallActiveDetection(structResponse)
	}

	if self.sleep != 0 {
		time.Sleep(time.Duration(self.sleep) * time.Second)
	} else if sleep := database.Memory.GetInt("HTTP Time Sleep"); sleep != 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	return structResponse, nil
}
