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

// Plugins ::
type Plugins struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
	wg      sync.WaitGroup
}

// Enumerate ::
func (options *Plugins) Enumerate() {
	plugins := func() map[string]bool {
		request, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(request.Body)

		if err != nil {
			printer.Fatal(err)
		}

		re := regexp.MustCompile("/wp-content/plugins/(.+?)/")

		submatchall := re.FindAllSubmatch([]byte(bodyBytes), -1)

		pluginsMapper := make(map[string]bool)

		for _, plugin := range submatchall {
			plugin := fmt.Sprintf("%s", plugin[1])

			pluginsMapper[plugin] = true
		}

		return pluginsMapper
	}()

	options.wg.Add(1)
	go func() {
		if len(plugins) > 0 {
			printer.Done("⎡ Plugin(s) :")
		}

		for plugin, _ := range plugins {
			printer.Done("⎢", plugin)

			if options.Verbose {
<<<<<<< Updated upstream
				printer.Warning("—", "URL Path:", options.HTTP.URL.Simple+"wp-content/plugins"+plugin)
=======
				printer.Warning("—", "URL Path:", options.HTTP.URL.Simple+"wp-content/plugins/"+plugin)
>>>>>>> Stashed changes
				options.readme(plugin)
				options.changelog(plugin)
				// options.license(plugin)
				// options.fullpathdisclosure(plugin)

			}

			time.Sleep(time.Millisecond)
		}

		defer options.wg.Done()
	}()

	options.wg.Wait()

	printer.Println("")
}

func (options *Plugins) fullpathdisclosure(plugin string) {
	for _, value := range wordlist.WPfpd {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

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

func (options *Plugins) readme(plugin string) {

	for _, value := range wordlist.WPreadme {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

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

func (options *Plugins) license(plugin string) {
	for _, value := range wordlist.WPlicense {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

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

func (options *Plugins) changelog(plugin string) {
	for _, value := range wordlist.WPchangesLogs {
		options.HTTP.URL.Directory = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

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
