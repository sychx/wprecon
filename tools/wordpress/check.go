package tools

import (
	. "fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/blkzy/wpsgo/pkg/gohttp"
	"github.com/blkzy/wpsgo/pkg/printer"
)

func Check(target string) (string, error) {
	var calc int
	var exists int
	var err error
	var response gohttp.Response
	var content []byte
	var payloads = [...]string{
		`<meta name="generator content="WordPress`,
		`<a href="http://www.wordpress.com">Powered by WordPress</a>`,
		`<link rel='https://api.w.org/'`}
	var paths = [...]string{
		"wp-content/uploads/",
		"wp-content/plugins/",
		"wp-content/themes/",
		"wp-includes/",
		"wp-admin/"}

	go func(URL string, htmlPayloads [3]string) {
		response, err = gohttp.HttpRequest(gohttp.Http{URL: URL})

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

	go func(URL string, directories [5]string) {
		for _, directory := range directories {
			request, err := gohttp.HttpRequest(gohttp.Http{URL: URL + directory})

			if err != nil {
				log.Fatal(err)
			}

			body, err := ioutil.ReadAll(request.Body)

			if err != nil {
				log.Fatal(err)
			}

			if directory == "wp-admin/" && request.StatusCode == 200 {
				printer.Warning("Apparently there is")
				exists++
			} else if !strings.Contains("Index Of", string(body)) {
				printer.Done("Index of Exists")
				exists++
			}
		}
	}(target, paths)

	calc = exists / 8 * 100

	return Sprintf("%v", calc), nil
}
