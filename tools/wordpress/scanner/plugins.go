package wpscan

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/text"
	"github.com/blackcrw/wprecon/pkg/wordlist"
)

type Plugins struct {
	Verbose bool
	Request gohttp.Http
}

func (options *Plugins) Changelog(plugin string) (bool, gohttp.Response) {
	for _, value := range wordlist.WPchangesLogs {
		options.Request.Dir = fmt.Sprintf("/wp-content/plugins/%s/%s", plugin, value)

		printer.Danger(options.Request.Dir)

		response, err := gohttp.HttpRequest(options.Request)

		if err != nil {
			printer.Fatal(err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		if response.StatusCode == 200 && string(bodyBytes) != "" {
			return true, response
		}
	}

	return false, gohttp.Response{}
}

func (options *Plugins) Enumerate() (bool, []string) {
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

	plugins := make([]string, len(submatchall))
	for key, plugin := range submatchall {
		pluginString := fmt.Sprintf("%s", plugin[1])

		if _, has := text.ContainsSliceString(plugins, pluginString); !has {
			plugins[key] = pluginString
		}
	}

	if len(plugins) != 0 {
		return true, plugins
	}

	return false, plugins
}
