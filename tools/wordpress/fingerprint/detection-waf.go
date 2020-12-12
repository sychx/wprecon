package wpfinger

import (
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// WAF :: This function will call the WAF's functions until one returns true.
// Yes, I could use anonymous functions, but I chose not to use it so that the code is not confused if you want to contribute in the future.
func WAF(options gohttp.Http) (bool, int, string) {
	printer.Loading("Active WAF detection module")

	if has, status, name := wordfence(options); has {

		return true, status, name
	} else if has, status, name := bulletproof(options); has {

		return true, status, name
	} else if has, status, name := betterwp(options); has {

		return true, status, name
	} else if has, status, name := sucuri(options); has {

		return true, status, name
	} else if has, status, name := wpsecurity(options); has {

		return true, status, name
	} else if has, status, name := allinonewpsecurity(options); has {

		return true, status, name
	} else if has, status, name := scanprotection(options); has {

		return true, status, name
	}

	return false, 0, ""
}

func wordfence(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/wordfence/"

	response, err := gohttp.HttpRequest(options)

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

func bulletproof(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/bulletproof-security/"

	response, err := gohttp.HttpRequest(options)

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

func betterwp(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/better-wp-security/"

	response, err := gohttp.HttpRequest(options)

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

func sucuri(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/sucuri-scanner/"

	response, err := gohttp.HttpRequest(options)

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

func wpsecurity(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/wp-security-scan/"

	response, err := gohttp.HttpRequest(options)

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

func allinonewpsecurity(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/all-in-one-wp-security-and-firewall/"

	response, err := gohttp.HttpRequest(options)

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

func scanprotection(options gohttp.Http) (bool, int, string) {
	options.Dir = "wp-content/plugins/6scan-protection/"

	response, err := gohttp.HttpRequest(options)

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
