package scanner

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"
	"time"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/wordlist"
)

// Themes ::
type Themes struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
	wg      sync.WaitGroup
}

// Enumerate ::
func (options *Themes) Enumerate() {
	themes := func() map[string]bool {
		request, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(request.Body)

		if err != nil {
			printer.Fatal(err)
		}

		re := regexp.MustCompile("/wp-content/themes/(.+?)/")

		submatchall := re.FindAllSubmatch([]byte(bodyBytes), -1)

		themesMapper := make(map[string]bool)

		for _, theme := range submatchall {
			theme := fmt.Sprintf("%s", theme[1])

			themesMapper[theme] = true
		}

		return themesMapper
	}()

	options.wg.Add(1)
	go func() {
		if len(themes) > 0 {
			printer.Done("⎡ Theme(s) :")
		}

		for theme, _ := range themes {
			printer.Done("⎢", theme, "—", options.HTTP.URL.Simple+"wp-content/themes/"+theme)

			if options.Verbose {
				options.readme(theme)
				options.changelog(theme)
				// options.license(theme)
				// options.fullpathdisclosure(theme)
			}

			time.Sleep(time.Millisecond)
		}

		defer options.wg.Done()
	}()

	options.wg.Wait()

	printer.Println("")
}

func (options *Themes) fullpathdisclosure(theme string) {
	for _, value := range wordlist.WPfpd {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/themes/%s/%s", theme, value)

		response, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 || response.StatusCode == 406 && string(bodyBytes) != "" {
			printer.Warning("— Full Path Disclosure:", response.URL.Full)

			break
		}
	}
}

func (options *Themes) readme(theme string) {

	for _, value := range wordlist.WPreadme {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/themes/%s/%s", theme, value)

		response, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 || response.StatusCode == 406 && string(bodyBytes) != "" {
			printer.Warning("— Readme:", response.URL.Full)

			break
		}
	}
}

func (options *Themes) license(theme string) {
	for _, value := range wordlist.WPlicense {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/themes/%s/%s", theme, value)

		response, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 || response.StatusCode == 406 && string(bodyBytes) != "" {
			printer.Warning("— License:", response.URL.Full)

			break
		}
	}
}

func (options *Themes) changelog(theme string) {
	for _, value := range wordlist.WPchangesLogs {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/themes/%s/%s", theme, value)

		response, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 || response.StatusCode == 406 && string(bodyBytes) != "" {
			printer.Warning("— Changelog:", response.URL.Full)

			break
		}
	}
}
