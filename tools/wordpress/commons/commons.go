package commons

import (
	"io/ioutil"
	"strings"

	"github.com/blackcrw/wprecon/pkg/gohttp"
)

func RobotsTXT(URL string, UserAgent bool, TOR bool, TLSCertificate bool) (bool, error) {
	http := gohttp.HTTPOptions{
		URL: gohttp.URLOptions{
			Simple:    URL,
			Directory: "robots.txt",
		},
		Options: gohttp.Options{
			RandomUserAgent:      UserAgent,
			Tor:                  TOR,
			TLSCertificateVerify: TLSCertificate,
		},
	}

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		return false, err
	}

	rawBytes, err := ioutil.ReadAll(response.Raw)

	if err != nil {
		return false, err
	}

	if response.StatusCode == 200 && string(rawBytes) != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func Sitemap(URL string, UserAgent bool, TOR bool, TLSCertificate bool) (bool, error) {
	http := gohttp.HTTPOptions{
		URL: gohttp.URLOptions{
			Simple:    URL,
			Directory: "sitemap.xml",
		},
		Options: gohttp.Options{
			RandomUserAgent:      UserAgent,
			Tor:                  TOR,
			TLSCertificateVerify: TLSCertificate,
		},
	}

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		return false, err
	}

	rawBytes, err := ioutil.ReadAll(response.Raw)

	if err != nil {
		return false, err
	}

	if response.StatusCode == 200 && string(rawBytes) != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func XMLRPC(URL string, UserAgent bool, TOR bool, TLSCertificate bool) (bool, error) {
	http := gohttp.HTTPOptions{
		URL: gohttp.URLOptions{
			Simple:    URL,
			Directory: "xmlrpc.php",
		},
		Options: gohttp.Options{
			RandomUserAgent:      UserAgent,
			Tor:                  TOR,
			TLSCertificateVerify: TLSCertificate,
		},
	}

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		return false, err
	}

	rawBytes, err := ioutil.ReadAll(response.Raw)

	if err != nil {
		return false, err
	}

	if response.StatusCode == 200 && string(rawBytes) != "" || strings.Contains(string(rawBytes), "XML-RPC server accepts POST requests only.") {
		return true, nil
	} else {
		return false, nil
	}
}
