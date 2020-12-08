package wpsfinger

import (
	"github.com/blackcrw/wpsgo/pkg/gohttp"
	"github.com/blackcrw/wpsgo/pkg/printer"
)

func WAF(target string) (bool, string) {
	printer.Loading("Active WAF detection module")

	if has, status, name := wordfence(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	} else if has, status, name := bulletproof(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	} else if has, status, name := betterwp(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	} else if has, status, name := sucuri(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	} else if has, status, name := wpsecurity(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	} else if has, status, name := allinonewpsecurity(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	} else if has, status, name := scanprotection(target); has {
		printer.LoadingWarning("WAF :", name, "Detected", "Status Code :", status)

		return has, name
	}

	printer.LoadingDanger("No WAF was detected! But that doesn't mean it doesn't.")
	return false, ""
}

func wordfence(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wordfence/"})

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

func bulletproof(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/bulletproof-security/"})

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

func betterwp(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/better-wp-security/"})

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

func sucuri(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/sucuri-scanner/"})

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

func wpsecurity(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wp-security-scan/"})

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

func allinonewpsecurity(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/all-in-one-wp-security-and-firewall/"})

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

func scanprotection(URL string) (bool, int, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/6scan-protection/"})

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
