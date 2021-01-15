package security

import (
	"fmt"
	"strings"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/tools/wordpress/extensions"
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
	printer.Done(":: Active WAF Agressive Detection Module ::")

	for _, path := range PathWAF {
		pathFormat := fmt.Sprintf("wp-content/plugins/%s/", path)
		response := extensions.SimpleRequest(InfosWprecon.Target, pathFormat)

		if response.StatusCode == 200 || response.StatusCode == 403 {
			printer.Done("I found this WAF")
			printer.Warning("Name \t:", strings.Title(strings.ReplaceAll(path, "-", " ")))
			printer.Warning("Status Code\t:", fmt.Sprint(response.StatusCode))
			printer.Warning("URL \t:", response.URL.Full)

			if importantfile := extensions.GetOneImportantFile(response.Raw); importantfile != "" {
				response2 := extensions.SimpleRequest(InfosWprecon.Target, pathFormat+importantfile)

				if readme := extensions.GetVersionStableTag(response2.Raw); readme != "" {
					printer.Warning("Version \t:", readme)
				} else if changelog := extensions.GetVersionChangelog(response2.Raw); changelog != "" {
					printer.Warning("Version \t:", changelog)
				}
			}

			scan := printer.ScanQ("Do you wish to continue ?! [Y/n] : ")

			if scan == "n" && scan != "\n" {
				printer.Fatal("Exiting...")
			}

			return response
		}
	}

	printer.Warning(":: No WAF was detected! But that doesn't mean it doesn't. ::")

	return nil
}
