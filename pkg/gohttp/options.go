package gohttp

import (
	"io"
	"net/http"
	"net/url"
)

// HTTPOptions :: This is Struct Http, it will inherit the struct Options and client.
type HTTPOptions struct {
	Method        string
	Proxy         func(*http.Request) (*url.URL, error)
	URL           URLOptions
	Options       Options
	TotalRequests int
}

// Response :: This struct will store the request data, and will be used for a return.
type Response struct {
	Method     string
	StatusCode int
	UserAgent  string
	Raw        io.Reader
	URL        URLOptions
}

// URLOptions :: This struct will be used to inform directories ... the complete URL ... or just the domain.
// (Alert) The focus of this struct is to be used together with HTTPOptions!
type URLOptions struct {
	Simple    string
	Full      string
	Directory string
}

// Options :: Here will be just a few "fetuares" that the user informed through the flags.
// (Alert) The focus of this struct is to be used together with HTTPOptions!
type Options struct {
	RandomUserAgent      bool
	TLSCertificateVerify bool
	Tor                  bool
}
