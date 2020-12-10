package wpfinger

import (
	"io/ioutil"
	"strings"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

func HasWordpress(target string, randomUserAgent bool) float32 {
	var exists float32
	var err error
	var response gohttp.Response
	var content []byte
	var payloads = [...]string{
		`<meta name="generator content="WordPress`,
		`<a href="http://www.wordpress.com">Powered by WordPress</a>`,
		`<link rel='https://api.w.org/'`}
	var directories = [...]string{
		"wp-content/uploads/",
		"wp-content/plugins/",
		"wp-content/themes/",
		"wp-includes/",
		"wp-admin/"}

	func(URL string, htmlPayloads [3]string) {
		response, err = gohttp.HttpRequest(gohttp.Http{URL: URL, RandomUserAgent: randomUserAgent})

		if err != nil {
			printer.Fatal(err)
		}

		content, err = ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		for _, value := range htmlPayloads {
			if strings.Contains(value, string(content)) {
				exists++
			}
		}

	}(target, payloads)

	for _, directory := range directories {
		func(URL string, directory string) {
			request, err := gohttp.HttpRequest(gohttp.Http{URL: URL, Dir: directory, RandomUserAgent: randomUserAgent})

			if err != nil {
				printer.Fatal(err)
			}

			body, err := ioutil.ReadAll(request.Body)

			if err != nil {
				printer.Fatal(err)
			}

			if directory == "wp-admin/" && request.StatusCode == 200 || request.StatusCode == 403 {
				printer.Warning("Status Code:", request.StatusCode, "—", "URL:", request.URLFULL)
				exists++
			} else if strings.Contains("Index Of", string(body)) {
				printer.Done("Listing enable:", request.URLFULL)
				exists++
			}

		}(target, directory)
	}

	return exists / 8 * 100
}
