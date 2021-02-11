package gohttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/blackbinn/wprecon/pkg/handler"
)

// TorURL ::
const (
	TorURL = "http://127.0.0.1:9080"
)

// Tor :: This will return the correctly formatted tor url.
func Tor() (func(*http.Request) (*url.URL, error), error) {
	tor, err := url.Parse(TorURL)

	if err != nil {
		return nil, fmt.Errorf("proxy URL is invalid (%w)", err)
	}

	return http.ProxyURL(tor), nil
}

// TorGetIP :: This will perform a check to see if your tor network is online or not.
func TorGetIP() string {
	defer handler.HandlerErrorTorProxy()

	http, _ := NewHTTPClient().SetURLFull("https://check.torproject.org/api/ip").SetSleep(0).OnTor(true)

	response, err := http.Run()

	if err != nil {
		panic(err)
	}

	var marshal map[string]interface{}

	err = json.Unmarshal([]byte(response.Raw), &marshal)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", marshal["IP"])
}
