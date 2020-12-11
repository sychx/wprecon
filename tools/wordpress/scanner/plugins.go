package wpscan

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/text"
)

func PluginsFind(options gohttp.Http) (bool, []string) {
	response, err := gohttp.HttpRequest(options)

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
