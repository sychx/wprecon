package net

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/blackcrw/wprecon/internal/printer"
)

// Tor :: This will return the correctly formatted tor url.
func Tor() (func(*http.Request) (*url.URL, error)) {
	var tor, _ = url.Parse("http://127.0.0.1:9080")

	return http.ProxyURL(tor)
}

// TorGetIP :: This will perform a check to see if your tor network is online or not.
func TorGetIP() string {
	var http = NewNETClient().SetURLFull("https://check.torproject.org/api/ip").SetSleep(0).OnTor(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Fatal(err)
	}
	
	var marshal map[string]interface{}
	
	err = json.Unmarshal([]byte(response.Raw), &marshal)
	
	if err != nil {
		printer.Fatal(err)
	}

	return fmt.Sprintf("%s", marshal["IP"])
}