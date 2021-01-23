package gohttp

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// HTTPOptions :: This is Struct Http, it will inherit the struct Options and client.
type httpoptions struct {
	url                  *URLOptions
	method               string
	tlsCertificateVerify bool
	tor                  bool
	proxy                func(*http.Request) (*url.URL, error)
	data                 io.Reader
	userAgent            string
	redirect             func(req *http.Request, via []*http.Request) error
	contentType          string
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

	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Fatal(err)
	}

	return response
}

func NewHTTPClient() *httpoptions {
	options := &httpoptions{
		method:      "GET",
		userAgent:   "WPrecon - Wordpress Recon (Vulnerability Scanner)",
		data:        nil,
		contentType: "text/html; charset=UTF-8"}

	options.url = &URLOptions{}

	return options
}

func (options *httpoptions) SetURL(url string) *httpoptions {
	if !strings.HasSuffix(url, "/") {
		options.url.Simple = url + "/"
		options.url.Full = url + "/"
	} else {
		options.url.Simple = url
		options.url.Full = url
	}

	return options
}

func (options *httpoptions) SetURLDirectory(directory string) *httpoptions {
	if !strings.HasPrefix(directory, "/") && !strings.HasSuffix(options.url.Simple, "/") {
		options.url.Directory = "/" + directory
		options.url.Full = options.url.Simple + "/" + directory
	} else {
		options.url.Directory = directory
		options.url.Full = options.url.Simple + directory
	}

	return options
}

func (options *httpoptions) SetURLFull(full string) *httpoptions {
	options.url.Full = full

	return options
}

func (options *httpoptions) OnTor(status bool) (*httpoptions, error) {
	if status {
		tor, err := url.Parse("http://127.0.0.1:9080")

		if err != nil {
			return nil, fmt.Errorf("proxy URL is invalid (%w)", err)
		}

		options.proxy = http.ProxyURL(tor)
	}

	return options, nil
}

func (options *httpoptions) OnRandomUserAgent(status bool) *httpoptions {
	if status {
		options.userAgent = randomuseragent()
	}

	return options
}

func (options *httpoptions) OnTLSCertificateVerify(status bool) *httpoptions {
	options.tlsCertificateVerify = status

	return options
}

func (options *httpoptions) SetMethod(method string) *httpoptions {
	options.method = method

	return options
}

func (options *httpoptions) SetUserAgent(useragent string) *httpoptions {
	options.userAgent = useragent

	return options
}

func (options *httpoptions) SetForm(form *url.Values) *httpoptions {
	options.data = strings.NewReader(form.Encode())

	return options
}

func (options *httpoptions) SetData(data string) *httpoptions {
	options.data = strings.NewReader(data)

	return options
}

func (options *httpoptions) SetRedirectFunc(redirectFunc func(req *http.Request, via []*http.Request) error) *httpoptions {
	options.redirect = redirectFunc

	return options
}

func (options *httpoptions) SetContentType(contenttype string) *httpoptions {
	options.contentType = contenttype

	return options
}

func (options *httpoptions) Run() (*Response, error) {
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

	InfosWprecon.TotalRequests++

	return &Response{
		Raw:      string(raw),
		URL:      options.url,
		Response: response,
	}, nil
}
