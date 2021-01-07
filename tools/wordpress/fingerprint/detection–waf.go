package fingerprint

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// WebApplicationFirewall ::
type WebApplicationFirewall struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
}

// Detection :: It is this function that must be performed for the detection of the web application firewall to be performed.
func (options *WebApplicationFirewall) Detection() {

	detection := func() (bool, int, string) {
		if has, status, name := options.wordfence(); has {
			return true, status, name
		} else if has, status, name := options.bulletproof(); has {

			return true, status, name
		} else if has, status, name := options.betterwp(); has {

			return true, status, name
		} else if has, status, name := options.sucuri(); has {

			return true, status, name
		} else if has, status, name := options.wpsecurity(); has {

			return true, status, name
		} else if has, status, name := options.allinonewpsecurity(); has {

			return true, status, name
		} else if has, status, name := options.scanprotection(); has {

			return true, status, name
		}

		return false, 0, ""
	}

	topline := printer.NewTopLine(":: Active WAF detection module ::")

	if has, status, name := detection(); has {
		topline.Warning(fmt.Sprint(status), "—", "WAF :", name)

		printer.Warning("Do you wish to continue ?! [Y/n] :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if strings.ToLower(scanner.Text()) != "y" {
			printer.Fatal("Exiting...")
		}
	} else {
		topline.Danger(":: No WAF was detected! But that doesn't mean it doesn't. ::")
	}

	printer.Println("")
}

func (options *WebApplicationFirewall) wordfence() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/wordfence/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Wordfence Security"
	}

	return false, response.StatusCode, ""
}

func (options *WebApplicationFirewall) bulletproof() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/bulletproof-security/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "BulletProof Security"
	}

	return false, response.StatusCode, ""
}

func (options *WebApplicationFirewall) betterwp() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/better-wp-security/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Better WP Security"
	}

	return false, response.StatusCode, ""
}

func (options *WebApplicationFirewall) sucuri() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/sucuri-scanner/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Sucuri Security"
	}

	return false, response.StatusCode, ""
}

func (options *WebApplicationFirewall) wpsecurity() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/wp-security-scan/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Acunetix WP Security"
	}

	return false, response.StatusCode, ""
}

func (options *WebApplicationFirewall) allinonewpsecurity() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/all-in-one-wp-security-and-firewall/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "All In One WP Security & Firewall"
	}

	return false, response.StatusCode, ""
}

func (options *WebApplicationFirewall) scanprotection() (bool, int, string) {
	options.HTTP.URL.Directory = "wp-content/plugins/6scan-protection/"

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "6Scan Security"
	}

	return false, response.StatusCode, ""
}
