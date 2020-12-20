package gohttp

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TorURL ::
const (
	TorURL = "http://127.0.0.1:9080"
)

type torAPIJSON struct {
	IsTor bool   `json:"IsTor"`
	IP    string `json:"IP"`
}

// Tor :: This will return the correctly formatted tor url.
func Tor() (func(*http.Request) (*url.URL, error), error) {
	tor, err := url.Parse(TorURL)

	if err != nil {
		return nil, fmt.Errorf("proxy URL is invalid (%w)", err)
	}

	return http.ProxyURL(tor), nil
}

// TorCheck :: This will perform a check to see if your tor network is online or not.
func TorCheck() (string, error) {
	var torJSON torAPIJSON

	proxy, err := Tor()

	if err != nil {
		return "0.0.0.0", err
	}

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
		return "0.0.0.0", err
	}

	resp, err := client.Do(request)

	if err != nil {
		return "0.0.0.0", err
	}

	if err := json.NewDecoder(resp.Body).Decode(&torJSON); err != nil {
		return "0.0.0.0", err
	}

	if !torJSON.IsTor {
		return torJSON.IP, fmt.Errorf("Proxy TOR not connected")
	}

	return torJSON.IP, nil
}
