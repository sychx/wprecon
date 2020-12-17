package gohttp

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// https://check.torproject.org/api/ip

type torAPIJSON struct {
	IsTor bool   `json:"IsTor"`
	IP    string `json:"IP"`
}

func TorCheck() (func(*http.Request) (*url.URL, error), error) {
	var torJSON torAPIJSON

	tor, err := url.Parse("http://127.0.0.1:9080")

	if err != nil {
		return nil, fmt.Errorf("proxy URL is invalid (%w)", err)
	}

	proxy := http.ProxyURL(tor)

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	request, err := http.NewRequest("GET", "https://check.torproject.org/api/ip", nil)

	if err != nil {
		return nil, fmt.Errorf("ProxyConnect tcp: dial tcp 127.0.0.1:9080: connect: connection refused")
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("ProxyConnect tcp: dial tcp 127.0.0.1:9080: connect: connection refused")
	}

	if err := json.NewDecoder(response.Body).Decode(&torJSON); err != nil {
		return nil, err
	}

	if !torJSON.IsTor {
		return nil, fmt.Errorf("Proxy TOR not connected")
	}

	return proxy, nil
}
