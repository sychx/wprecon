package wpfinger

import (
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

type WAF struct {
	Request gohttp.Http
	Verbose bool
}

// WAF :: This function will call the WAF's functions until one returns true.
// Yes, I could use anonymous functions, but I chose not to use it so that the code is not confused if you want to contribute in the future.
func (options *WAF) Detection() (bool, int, string) {

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

func (options *WAF) wordfence() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/wordfence/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Wordfence Security"
	}

	return false, response.StatusCode, ""
}

func (options *WAF) bulletproof() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/bulletproof-security/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "BulletProof Security"
	}

	return false, response.StatusCode, ""
}

func (options *WAF) betterwp() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/better-wp-security/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Better WP Security"
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "Better WP Security"
	case 403:
		return true, 403, "Better WP Security"
	}

	return false, response.StatusCode, ""
}

func (options *WAF) sucuri() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/sucuri-scanner/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Sucuri Security"
	}

	return false, response.StatusCode, ""
}

func (options *WAF) wpsecurity() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/wp-security-scan/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "Acunetix WP Security"
	}

	return false, response.StatusCode, ""
}

func (options *WAF) allinonewpsecurity() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/all-in-one-wp-security-and-firewall/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "All In One WP Security & Firewall"
	}

	return false, response.StatusCode, ""
}

func (options *WAF) scanprotection() (bool, int, string) {
	options.Request.Dir = "wp-content/plugins/6scan-protection/"

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		return true, response.StatusCode, "6Scan Security"
	}

	return false, response.StatusCode, ""
}
