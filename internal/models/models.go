package models

import (
	"io"
	"net/http"
	"net/url"
)

type InformationsModel struct {
	Url string
	Status int
	Raw string
	Confidence int
	FoundBy string
}

type VersionApiModel struct {
	Configure struct {
		Version string `json:"version"`
	} `json:"Configure"`
}

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
