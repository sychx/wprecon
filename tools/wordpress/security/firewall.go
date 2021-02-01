package security

import (
	"fmt"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
)

// WAFAggressiveDetection :: It is this function that must be performed for the detection of the web application firewall to be performed.
func WAFAggressiveDetection() *gohttp.Response {
	var pathWAF = [...]string{
		"wordfence",
		"cloudflare",
		"bulletproof-security",
		"better-wp-security",
		"sucuri-scanner",
		"wp-security-scan",
		"block-bad-queries",
		"all-in-one-wp-security-and-firewall",
		"6scan-protection",
		"siteguard",
		"ninjafirewall",
		"malcare-security",
		"wp-cerber",
		"wesecur-security"}

	for _, path := range pathWAF {
		pathFormat := fmt.Sprintf("%s/plugins/%s/", InfosWprecon.WPContent, path)

		if response := gohttp.SimpleRequest(InfosWprecon.Target, pathFormat); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
			return response
		}
	}

	return nil
}
