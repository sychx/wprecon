package fingerprint

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// Honeypot ::
type Honeypot struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
}

// Detection ::
func (options *Honeypot) Detection() {
	var question string

	host, err := gohttp.GetHost(options.HTTP.URL.Simple)

	if err != nil {
		printer.Fatal(err)
	}

	ips, err := net.LookupIP(host)

	if err != nil {
		printer.Fatal(err)
	}

	optionshttp := gohttp.HTTPOptions{
		URL: gohttp.URLOptions{
			Simple:    "https://api.shodan.io",
			Directory: fmt.Sprintf("/labs/honeyscore/%s?key=C23OXE0bVMrul2YeqcL7zxb6jZ4pj2by", ips[0]),
		},
		Options: gohttp.Options{
			RandomUserAgent:      true,
			TLSCertificateVerify: false,
		},
	}

	request, err := gohttp.HTTPRequest(&optionshttp)

	if err != nil {
		printer.Fatal(err)
	}

	body, err := ioutil.ReadAll(request.Raw)

	if err != nil {
		printer.Fatal(err)
	}

	x, err := strconv.ParseFloat(string(body), 32)

	if err != nil {
		printer.Fatal(err)
	}

	convert := options.convertfloat(string(body))

	if x > 0.7 {
		printer.Danger("With a", convert, "chance of this host being a Honeypot. Do you wish to continue ?! [y/N] ")
		fmt.Scan(&question)

		if strings.ToLower(question) != "y" {
			printer.Fatal("Exiting...")
		}
	} else {
		printer.Done("With a", convert, "chance of this host being a Honeypot.")
	}
}

func (options *Honeypot) convertfloat(text string) string {
	strings.ReplaceAll(text, "0.0", "0%")
	strings.ReplaceAll(text, "0.1", "10%")
	strings.ReplaceAll(text, "0.2", "20%")
	strings.ReplaceAll(text, "0.3", "30%")
	strings.ReplaceAll(text, "0.4", "40%")
	strings.ReplaceAll(text, "0.5", "50%")
	strings.ReplaceAll(text, "0.6", "60%")
	strings.ReplaceAll(text, "0.7", "70%")
	strings.ReplaceAll(text, "0.8", "80%")
	strings.ReplaceAll(text, "0.9", "90%")
	strings.ReplaceAll(text, "1.0", "100%")

	return text
}
