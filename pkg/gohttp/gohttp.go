package gohttp

import (
	"io"
)

// Http :: This is Struct Http, it will inherit the struct Options and client.
type Http struct {
	Method               string
	URL                  string
	URLFULL              string
	Dir                  string
	Proxy                string
	RandomUserAgent      bool
	TLSCertificateVerify bool
}

// Options :: This struct will keep the options that can be used, random agent among others.
type Options struct {
	RandomUserAgent bool
}

// Response :: This struct will store the request data, and will be used for a return.
type Response struct {
	URL        string
	URLFULL    string
	Dir        string
	Method     string
	StatusCode int
	UserAgent  string
	Body       io.Reader
}
