package models

import (
	"io"
	"net/http"
	"net/url"
)

type EnumerateModel struct{
	Name       string
	FoundBy    string
	Others     []EnumerateOthersModel
}

type EnumerateOthersModel struct {
	Confidence int8
	Version    string
	FoundBy    string
	Match      []string	
}

type InterestingModel struct {
	Url        string
	Status     int
	Raw        string
	Confidence int8
	FoundBy    string
}

type FindingsVersionModel struct {
	Version string
	Match string
}

type ConfigModel struct {
	App struct {
		Name        string `yaml:"name"`
		Author      string `yaml:"author"`
		Description string `yaml:"description"`
		Version     string `yaml:"version"`
		ApiUrl      string `yaml:"api_url"`
	} `yaml:"application"`
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

type GetVersions interface {
	GetVersionByStableTag(raw string)  []string
	GetVersionByChangelog(raw string)  []string
	GetVersionByReleaseLog(raw string) []string
}
