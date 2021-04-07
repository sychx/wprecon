package enumerate

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/blackbinn/wprecon/internal/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
)

const (
	MatchPluginPassive            = "/plugins/(.*?)/.*?[css|js].*?ver=(\\d{1,2}\\.\\d{1,2}\\.\\d{1,3})"
	MatchPluginAgressiveDirectory = "<a href=\"(.*?)/\">.*?/</a>"
	MatchPluginAgressiveNoVersion = "/plugins/(.*?)/.*?[.css|.js]"
)

type plugin struct {
	raw           string
	wpContentPath string
	plugins       [][]string
}

func NewPlugins(raw string, wpcontent string) *plugin {
	return &plugin{raw: raw, wpContentPath: wpcontent}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (object *plugin) Passive() [][]string {
	var re = regexp.MustCompile(object.wpContentPath + MatchPluginPassive)

	return re.FindAllStringSubmatch(object.raw, -1)
}

func (object *plugin) Aggressive() [][]string {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		if response := commons.DirectoryPlugins(); response.Response.StatusCode == 200 {
			var re = regexp.MustCompile(MatchPluginAgressiveDirectory)

			for _, submatch := range re.FindAllStringSubmatch(response.Raw, -1) {
				/*
				   If the condition is true, it means that the plugin exists. Soon he will not add the list, but the condition will go to the else that will add another match for the plugins.
				   Note: For you to understand better, I recommend that you see this in operation.
				*/
				if _, contains := text.FindStringInSliceSlice(object.plugins, 1, submatch[1]); !contains {
					object.plugins = append(object.plugins, submatch)
				} else {
					var index = text.FindByValueInIndex(object.plugins, submatch[1])
					object.plugins[index][0] = fmt.Sprintf("%s,%s", object.plugins[index][0], submatch[0])
				}
			}
		}

		defer wg.Done()
	}()

	go func() {
		for _, submatch := range object.Passive() {
			if _, has := text.ContainsInSliceSlice(object.plugins, submatch[1]); !has {
				object.plugins = append(object.plugins, submatch)
			}
		}

		defer wg.Done()
	}()

	wg.Wait()

	return object.plugins
}
