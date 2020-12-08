package wpsfinger

import (
	"github.com/blackcrw/wpsgo/pkg/gohttp"
	"github.com/blackcrw/wpsgo/pkg/printer"
)

func WAF(target string) (bool, string) {

	if hasWaf, nameWaf := wordfence(target); hasWaf {
		return hasWaf, nameWaf
	} else if hasWaf, nameWaf := bulletproof(target); hasWaf {
		return hasWaf, nameWaf
	} else if hasWaf, nameWaf := betterwp(target); hasWaf {
		return hasWaf, nameWaf
	} else if hasWaf, nameWaf := sucuri(target); hasWaf {
		return hasWaf, nameWaf
	} else if hasWaf, nameWaf := wpsecurity(target); hasWaf {
		return hasWaf, nameWaf
	} else if hasWaf, nameWaf := allinonewpsecurity(target); hasWaf {
		return hasWaf, nameWaf
	} else if hasWaf, nameWaf := scanprotection(target); hasWaf {
		return hasWaf, nameWaf
	}

	return false, ""
}

func wordfence(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wordfence/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — Wordfence Security"
	case 403:
		return true, "403 — Wordfence Security"
	}
}

func bulletproof(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/bulletproof-security/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — BulletProof Security"
	case 403:
		return true, "403 — BulletProof Security"
	}
}

func betterwp(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/better-wp-security/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — Better WP Security"
	case 403:
		return true, "403 — Better WP Security"
	}
}

func sucuri(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/sucuri-scanner/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — Sucuri Security"
	case 403:
		return true, "403 — Sucuri Security"
	}
}

func wpsecurity(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wp-security-scan/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — Acunetix WP Security"
	case 403:
		return true, "403 — Acunetix WP Security"
	}
}

func allinonewpsecurity(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/all-in-one-wp-security-and-firewall/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — All In One WP Security & Firewall"
	case 403:
		return true, "403 — All In One WP Security & Firewall"
	}
}

func scanprotection(target string) (bool, string) {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/6scan-protection/"})

	if err != nil {
		printer.Fatal(err)
	}

	switch response.StatusCode {
	case 200:
		return true, "200 — 6Scan Security"
	case 403:
		return true, "403 — 6Scan Security"
	}

}
