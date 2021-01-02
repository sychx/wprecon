package gohttp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
)

// HTTPRequest :: This function will be used for any request that is made.
func HTTPRequest(options *HTTPOptions) (Response, error) {
	if !strings.HasSuffix(options.URL.Simple, "/") {
		options.URL.Simple = fmt.Sprintf("%s/", options.URL.Simple)
	}

	options.URL.Full = options.URL.Simple + options.URL.Directory

	if options.Method == "" {
		options.Method = "GET"
	}

	if options.Options.Tor {
		t, err := Tor()

		if err != nil {
			return Response{}, err
		}

		options.Proxy = t
	} else {
		options.Proxy = http.ProxyFromEnvironment
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: options.Proxy,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: options.Options.TLSCertificateVerify,
			},
		},
	}

	request, err := http.NewRequest(options.Method, options.URL.Full, nil)

	if err != nil {
		return Response{}, err
	}

	if options.Options.RandomUserAgent {
		request.Header.Set("User-Agent", randomuseragent())
	} else {
		request.Header.Set("User-Agent", "WPrecon - Wordpress Recon (Vulnerability Scanner) (GoHttp 1.0)")
	}

	resp, err := client.Do(request)

	if err != nil {
		return Response{}, err
	}

	httpResponse := Response{
		Method:     options.Method,
		URL:        options.URL,
		StatusCode: resp.StatusCode,
		UserAgent:  request.UserAgent(),
		Raw:        resp.Body,
	}

	options.TotalRequests++

	return httpResponse, nil
}
