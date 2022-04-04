package net

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net/middleware"
	"github.com/blackcrw/wprecon/internal/printer"
)

// net_options :: This is Struct Http, it will inherit the struct Options and client.
type net_options struct {
	tor                    bool
	tls_certificate_verify bool
	method                 string
	user_agent             string
	content_type           string
	data                   io.Reader
	url                    *models.UrlOptionsModel
	sleep                  time.Duration
	proxy                  func(*http.Request) (*url.URL, error)
	redirect               func(req *http.Request, via []*http.Request) error
}

// SimpleRequest :: The first parameter must always be the base url. The second must be the directory.
func SimpleRequest(url string) *models.ResponseModel {
	var http = NewNETClient()
	http.SetURLFull(url)
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil {
		printer.Fatal(err)
	}

	return response
}

func NewNETClient() *net_options {
	var net = &net_options{
		method:       "GET",
		user_agent:   "WPrecon - Wordpress Recon (Vulnerability Recon)",
		data:         nil,
		content_type: "text/html; charset=UTF-8"}

	net.url = &models.UrlOptionsModel{}

	return net
}

func (this *net_options) SetURL(url string) *net_options {
	if !strings.HasSuffix(url, "/") {
		this.url.Simple = url + "/"
		this.url.Full = url + "/"
	} else {
		this.url.Simple = url
		this.url.Full = url
	}

	return this
}

func (this *net_options) SetURLDirectory(directory string) *net_options {
	if !strings.HasPrefix(directory, "/") && !strings.HasSuffix(this.url.Simple, "/") {
		this.url.Directory = "/" + directory
		this.url.Full = this.url.Simple + "/" + directory
	} else {
		this.url.Directory = directory
		this.url.Full = this.url.Simple + directory
	}

	return this
}

func (this *net_options) SetURLFull(full string) *net_options {
	this.url.Full = full

	return this
}

func (this *net_options) OnTor(status bool) (*net_options) {
	if status {
		this.proxy = Tor()
	}

	return this
}

func (this *net_options) OnRandomUserAgent(status bool) *net_options {
	if status {
		this.user_agent = random_user_agent()
	}

	return this
}

func (this *net_options) OnTLSCertificateVerify(status bool) *net_options {
	this.tls_certificate_verify = status

	return this
}

func (this *net_options) SetMethod(method string) *net_options {
	this.method = method

	return this
}

func (this *net_options) SetUserAgent(userAgent string) *net_options {
	this.user_agent = userAgent

	return this
}

func (this *net_options) SetForm(form *url.Values) *net_options {
	this.data = strings.NewReader(form.Encode())

	return this
}

func (this *net_options) SetData(data string) *net_options {
	this.data = strings.NewReader(data)

	return this
}

func (this *net_options) SetRedirectFunc(redirectFunc func(req *http.Request, via []*http.Request) error) *net_options {
	this.redirect = redirectFunc

	return this
}

func (this *net_options) SetContentType(contentType string) *net_options {
	this.content_type = contentType

	return this
}

func (this *net_options) SetSleep(tm int) *net_options {
	this.sleep = time.Duration(tm)

	return this
}

func (this *net_options) Runner() (*models.ResponseModel, error) {
	var client = &http.Client{
		CheckRedirect: this.redirect,
		Transport: &http.Transport{
			Proxy:             this.proxy,
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: this.tls_certificate_verify,
			},
		},
	}

	var request, err = http.NewRequest(this.method, this.url.Full, this.data)
	if err != nil { return nil, err }

	request.Header.Set("User-Agent", this.user_agent)
	request.Header.Set("Content-Type", this.content_type)

	response, err := client.Do(request)
	if err != nil { return nil, err }

	raw, err := ioutil.ReadAll(response.Body)
	if err != nil { return nil, err }

	database.Memory.AddInt("HTTP Total")

	var struct_response = &models.ResponseModel{
		Raw:      string(raw),
		URL:      this.url,
		Response: response,
	}

	middleware.ActiveWebApplicationFirewall(struct_response)

	if sleep := database.Memory.GetInt("HTTP Time Sleep"); sleep != 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	return struct_response, nil
}
