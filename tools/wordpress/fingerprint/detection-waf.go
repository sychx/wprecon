package wpfinger

import (
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

func WAF(target string, randomUserAgent bool) (bool, string) {
	printer.Loading("Active WAF detection module")

	if has, status, name := wordfence(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	} else if has, status, name := bulletproof(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	} else if has, status, name := betterwp(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	} else if has, status, name := sucuri(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	} else if has, status, name := wpsecurity(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	} else if has, status, name := allinonewpsecurity(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	} else if has, status, name := scanprotection(target, randomUserAgent); has {
		printer.LoadingWarning("Status Code:", status, "—", "WAF:", name)

		return has, name
	}

	printer.LoadingDanger("No WAF was detected! But that doesn't mean it doesn't.")
	return false, ""
}

func wordfence(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wordfence/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "Wordfence Security"
	case 403:
		return true, 403, "Wordfence Security"
	}

	return false, 0, ""
}

func bulletproof(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/bulletproof-security/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "BulletProof Security"
	case 403:
		return true, 403, "BulletProof Security"
	}

	return false, 0, ""
}

func betterwp(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/better-wp-security/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "Better WP Security"
	case 403:
		return true, 403, "Better WP Security"
	}

	return false, 0, ""
}

func sucuri(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/sucuri-scanner/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "Sucuri Security"
	case 403:
		return true, 403, "Sucuri Security"
	}

	return false, 0, ""
}

func wpsecurity(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wp-security-scan/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "Acunetix WP Security"
	case 403:
		return true, 403, "Acunetix WP Security"
	}

	return false, 0, ""
}

func allinonewpsecurity(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/all-in-one-wp-security-and-firewall/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "All In One WP Security & Firewall"
	case 403:
		return true, 403, "All In One WP Security & Firewall"
	}

	return false, 0, ""
}

func scanprotection(URL string, randomUserAgent bool) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/6scan-protection/", Options: gohttp.Options{RandomUserAgent: randomUserAgent}})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, 200, "6Scan Security"
	case 403:
		return true, 403, "6Scan Security"
	}

	return false, 0, ""
}
