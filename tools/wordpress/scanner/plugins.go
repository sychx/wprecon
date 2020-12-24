package scanner

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// Plugins ::
type Plugins struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
}

// Enumerate ::
func (options *Plugins) Enumerate() {
	var pluginsMapper = make(map[string]bool)

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

	for _, plugin := range submatchall {
		plugin := fmt.Sprintf("%s", plugin[1])

		pluginsMapper[plugin] = true
	}

	if len(pluginsMapper) > 0 {
		printer.Done("⎡ Plugin(s) :")
	}

	for plugin, _ := range pluginsMapper {
		printer.Done("⎢", plugin)

		if options.Verbose {
			printer.Warning("—", "URL Path:", options.HTTP.URL.Simple+"wp-content/plugins/"+plugin)
		}
	}

	printer.Println("")
}
