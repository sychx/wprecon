package wpscan

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

type Plugins struct {
	Verbose bool
	Request gohttp.Http
}

func (options *Plugins) Enumerate() {
	wg := new(sync.WaitGroup)

	response, err := gohttp.HttpRequest(options.Request)

	if err != nil {
		printer.Fatal(err)
	}

	re := regexp.MustCompile("/wp-content/plugins/(.+?)/")

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		printer.Fatal(err)
	}

	submatchall := re.FindAllSubmatch([]byte(bodyBytes), -1)

	set := make(map[string]int)

	for _, plugin := range submatchall {
		plugin := fmt.Sprintf("%s", plugin[1])

		set[plugin] += 1

		if value, _ := set[plugin]; value <= 1 {
			printer.Done("Plugin:", plugin)

			if options.Verbose {
				wg.Add(4)

				go options.fullpathdisclosure(wg, plugin)
				go options.changelog(wg, plugin)
				go options.license(wg, plugin)
				go options.readme(wg, plugin)

				wg.Wait()

			}
		}
	}
}

func (options *Plugins) fullpathdisclosure(wg *sync.WaitGroup, plugin string) {
	defer wg.Done()

	for _, value := range wordlist.WPfpd {
		options.Request.Dir = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

		response, err := gohttp.HttpRequest(options.Request)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 && string(bodyBytes) != "" {
			printer.Warning("Full Path Disclosure:", response.URLFULL)

			break
		}
	}

	time.Sleep(time.Millisecond)
}

func (options *Plugins) readme(wg *sync.WaitGroup, plugin string) {
	defer wg.Done()

	for _, value := range wordlist.WPreadme {
		options.Request.Dir = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

		response, err := gohttp.HttpRequest(options.Request)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 && string(bodyBytes) != "" {
			printer.Warning("Readme:", response.URLFULL)

			break
		}

	}
	time.Sleep(time.Millisecond)
}

func (options *Plugins) license(wg *sync.WaitGroup, plugin string) {
	defer wg.Done()

	for _, value := range wordlist.WPlicense {
		options.Request.Dir = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

		response, err := gohttp.HttpRequest(options.Request)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 && string(bodyBytes) != "" {
			printer.Warning("License:", response.URLFULL)

			break
		}

	}
	time.Sleep(time.Millisecond)
}

func (options *Plugins) changelog(wg *sync.WaitGroup, plugin string) {
	defer wg.Done()

	for _, value := range wordlist.WPchangesLogs {
		options.Request.Dir = fmt.Sprintf("wp-content/plugins/%s/%s", plugin, value)

		response, err := gohttp.HttpRequest(options.Request)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 && string(bodyBytes) != "" {
			printer.Warning("Changelog:", response.URLFULL)

			break
		}

	}

	time.Sleep(time.Millisecond)
}
