package wpsfinger

import (
	"github.com/blkzy/wpsgo/pkg/gohttp"
	"github.com/blkzy/wpsgo/pkg/printer"
	"github.com/blkzy/wpsgo/tools/wps"
)

func WAF(target string) (bool, interface{}) {
	var WAF [6]string
	var hasWAF bool = false

	wps.Sync.Add(6)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wordfence/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — Wordfence Security"
			hasWAF = true
		case 403:
			WAF[0] = "403 — Wordfence Security"
		}

		wps.Sync.Done()
	}(target)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/bulletproof-security/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — BulletProof Security"
			hasWAF = true
		case 403:
			WAF[0] = "403 — BulletProof Security"
		}

		wps.Sync.Done()
	}(target)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/better-wp-security/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — Better WP Security"
			hasWAF = true
		case 403:
			WAF[0] = "403 — Better WP Security"
		}

		wps.Sync.Done()
	}(target)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/sucuri-scanner/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — Sucuri Security"
			hasWAF = true
		case 403:
			WAF[0] = "403 — Sucuri Security"
		}

		wps.Sync.Done()
	}(target)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/wp-security-scan/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — Acunetix WP Security"
			hasWAF = true
		case 403:
			WAF[0] = "403 — Acunetix WP Security"
		}

		wps.Sync.Done()
	}(target)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/all-in-one-wp-security-and-firewall/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — All In One WP Security & Firewall"
			hasWAF = true
		case 403:
			WAF[0] = "403 — All In One WP Security & Firewall"
		}

		wps.Sync.Done()
	}(target)

	go func(URL string) {
		response, err := gohttp.HttpRequest(gohttp.Http{URL: URL + "/wp-content/plugins/6scan-protection/"})

		if err != nil {
			printer.Fatal(err)
		}

		switch response.StatusCode {
		case 200:
			WAF[0] = "200 — 6Scan Security"
			hasWAF = true
		case 403:
			WAF[0] = "403 — 6Scan Security"
		}

		wps.Sync.Done()
	}(target)

	wps.Sync.Wait()

	return hasWAF, WAF
}
