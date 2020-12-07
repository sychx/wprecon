package gohttp

import (
	"io"
)

// Http :: This is Struct Http, it will inherit the struct Options and client.
type Http struct {
	Method  string
	URL     string
	Proxy   string
	Options Options
}

// Options :: This struct will keep the options that can be used, random agent among others.
type Options struct {
	RandomUserAgent bool
}

// Result :: This struct will store the request data, and will be used for a return.
type Result struct {
	URL        string
	Method     string
	StatusCode int64
	UserAgent  string
	Body       io.Reader
}
