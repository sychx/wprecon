package gohttp

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/blackcrw/wprecon/cli/config"
)

// HTTPRequest :: This function will be used for any request that is made.
func HTTPRequest(options *HTTPOptions) (Response, error) {
	var redirect string
	var data io.Reader

	if !strings.HasSuffix(options.URL.Simple, "/") {
		options.URL.Simple = fmt.Sprintf("%s/", options.URL.Simple)
	}

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

	if options.Data != "" {
		data = bytes.NewBuffer([]byte(options.Data))
		options.Method = "POST"
	} else {
		data = nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: options.Proxy,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: options.Options.TLSCertificateVerify,
			},
		},
	}

	options.URL.Full = options.URL.Simple + options.URL.Directory

	request, err := http.NewRequest(options.Method, options.URL.Full, data)

	if err != nil {
		return Response{}, err
	}

	if options.Options.RandomUserAgent {
		request.Header.Set("User-Agent", randomuseragent())
	} else {
		request.Header.Set("User-Agent", "WPrecon - Wordpress Recon (Vulnerability Scanner)")
	}

	resp, err := client.Do(request)

	if strings.Contains(fmt.Sprintf("%s", err), "proxyconnect tcp: dial tcp 127.0.0.1:9080: connect: connection refused") && options.Options.Tor {
		return Response{}, fmt.Errorf("Connection refused, the tor with the command: tor --HTTPTunnelPort 9080")
	}

	if strings.Contains(fmt.Sprintf("%s", err), "dial tcp 144.217.235.104:8777: connect: connection refused") && options.Options.Tor {
		return Response{}, fmt.Errorf("Connection refused to API")
	}

	if err != nil {
		return Response{}, err
	}

	raw, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Response{}, err
	}

	if resp.StatusCode == 302 {
		redirect = resp.Header.Get("Location")
	}

	httpResponse := Response{
		Method:      options.Method,
		URL:         options.URL,
		StatusCode:  resp.StatusCode,
		UserAgent:   request.UserAgent(),
		Raw:         string(raw),
		RawIo:       resp.Body,
		RedirectURL: redirect,
	}

	InfosWprecon.TotalRequests++

	return httpResponse, nil
}
