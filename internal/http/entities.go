package http

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type _EntityOptionsHttp struct {
	url 				 string
	data                 io.Reader
	method               string
	userAgent            string
	timeSleep 			 time.Duration
	contentType 		 string
	tlsCertificateVerify bool
}

type EntityResponse struct {
	RawIo    io.Reader
	Raw      string
	URL      *url.URL
	Response *http.Response
}