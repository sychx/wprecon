package wpscan

import (
	"fmt"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/text"
	"io/ioutil"
	"regexp"
)

func PluginEnum(target string, randomUserAgent bool) []string {
	response, err := gohttp.HttpRequest(gohttp.Http{URL: target, RandomUserAgent: randomUserAgent})

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

	return plugins
}

/*
func Changelogs(target string, randomUserAgent bool) {
	for _, file := range wordlist.WPchangesLogs {

		response, err := gohttp.HttpRequest(gohttp.Http{URL: target, Dir: "/wp-content/plugins/"+file, RandomUserAgent: randomUserAgent})

		if err != nil {
			printer.Fatal(err)
		}

		printer.Danger(response)
	}
}
*/