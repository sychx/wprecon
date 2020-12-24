package scanner

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// Themes ::
type Themes struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
	wg      sync.WaitGroup
}

// Enumerate ::
func (options *Themes) Enumerate() {
	var themesMapper = make(map[string]bool)

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

	for _, theme := range submatchall {
		theme := fmt.Sprintf("%s", theme[1])

		themesMapper[theme] = true
	}

	if len(themesMapper) > 0 {
		printer.Done("⎡ Theme(s) :")
	}

	for theme, _ := range themesMapper {
		printer.Done("⎢", theme)

		if options.Verbose {
			printer.Warning("—", "URL Path:", options.HTTP.URL.Simple+"wp-content/themes/"+theme)
		}
	}

	printer.Println("")
}
