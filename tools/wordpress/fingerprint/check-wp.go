package wpsfinger

import (
	. "fmt"
	"io/ioutil"
	"strings"

	"github.com/blkzy/wpsgo/pkg/gohttp"
	"github.com/blkzy/wpsgo/pkg/printer"
	"github.com/blkzy/wpsgo/tools/wps"
)

func HasWordpress(target string) string {
	var calc int
	var exists int
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

	wps.Sync.Add(2)

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

		wps.Sync.Done()
	}(target, payloads)

	for _, directory := range directories {
		wps.Sync.Add(1)

		go func(URL string, directory string) {
			request, err := gohttp.HttpRequest(gohttp.Http{URL: URL + directory})

			if err != nil {
				printer.Fatal(err)
			}

			body, err := ioutil.ReadAll(request.Body)

			if err != nil {
				printer.Fatal(err)
			}

			if directory == "wp-admin/" && request.StatusCode == 200 || request.StatusCode == 403 {
				printer.Warning("Status Code:", request.StatusCode, "in the URL:", URL+directory)
				exists++
			} else if !strings.Contains("Index Of", string(body)) {
				printer.Done("Listing enable:", URL+directory)
				exists++
			}

			defer wps.Sync.Done()
		}(target, directory)
	}

	wps.Sync.Wait()

	calc = exists / 8 * 100

	return Sprintf("%v", calc)
}
