package wpfinger

import (
	"io/ioutil"
	"strings"
	"sync"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

type Wordpress struct {
	Request  gohttp.Http
	Verbose  bool
	accuracy float32
}

func (options *Wordpress) Check() float32 {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go options.directory(wg)
	go options.htmlcode(wg)

	wg.Wait()

	return options.accuracy / 8 * 100
}

func (options *Wordpress) htmlcode(wg *sync.WaitGroup) {
	defer wg.Done()

	var payloads = [...]string{
		`<meta name="generator content="WordPress`,
		`<a href="http://www.wordpress.com">Powered by WordPress</a>`,
		`<link rel='https://api.w.org/'`}

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		printer.Fatal(err)
	}

	for _, value := range payloads {
		if strings.Contains(value, string(content)) {
			options.accuracy++
		}
	}
}

func (options *Wordpress) directory(wg *sync.WaitGroup) {
	defer wg.Done()

	var directories = [...]string{
		"wp-content/uploads/",
		"wp-content/plugins/",
		"wp-content/themes/",
		"wp-includes/",
		"wp-admin/"}

	for _, directory := range directories {
		wg.Add(1)
		go func(directory string) {
			defer wg.Done()

			options.Request.Dir = directory

			request, err := gohttp.HttpRequest(options.Request)

			if err != nil {
				printer.Fatal(err)
			}

			body, err := ioutil.ReadAll(request.Body)

			if err != nil {
				printer.Fatal(err)
			}

			if directory == "wp-admin/" && request.StatusCode == 200 || request.StatusCode == 403 {
				printer.Warning("Status Code:", request.StatusCode, "—", "URL:", request.URLFULL)
				options.accuracy++
			} else if strings.Contains("Index Of", string(body)) {
				printer.Done("Listing enable:", request.URLFULL)
				options.accuracy++
			}
		}(directory)
	}
}
