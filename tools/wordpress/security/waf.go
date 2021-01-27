package security

import (
	"fmt"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
)

var PathWAF = [...]string{
	"wordfence",
	"bulletproof-security",
	"better-wp-security",
	"sucuri-scanner",
	"wp-security-scan",
	"all-in-one-wp-security-and-firewall",
	"6scan-protection"}

// WAFAgressiveDetection :: It is this function that must be performed for the detection of the web application firewall to be performed.
func WAFAgressiveDetection() *gohttp.Response {
	for _, path := range PathWAF {
		pathFormat := fmt.Sprintf("wp-content/plugins/%s/", path)

		if response := gohttp.SimpleRequest(InfosWprecon.Target, pathFormat); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
			return response
		}
	}

	return nil
}
