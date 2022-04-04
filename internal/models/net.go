package models

import (
	"io"
	"net/http"
	"net/url"
)

type UrlOptionsModel struct {
	Simple    string
	Full      string
	Directory string
	URL       *url.URL
}

type ResponseModel struct {
	RawIo    io.Reader
	Raw      string
	URL      *UrlOptionsModel
	Response *http.Response
}

type MiddlewareFirewallModel struct {
	Name       string
	Solve      string
	Confidence int
	FoundBy    string
}