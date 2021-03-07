package enumerate

import (
	"regexp"
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
)

type plugin struct {
	raw     string
	plugins [][]string
}

func NewPlugins() *plugin {
	return &plugin{raw: database.Memory.GetString("HTTP Index Raw")}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (options *plugin) Passive() [][]string {
	regxp := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/plugins/(.*?)/.*?[css|js].*?ver=(\\d{1,2}\\.\\d{1,2}\\.\\d{1,3})")

	for _, plugin := range regxp.FindAllStringSubmatch(options.raw, -1) {
		formOfMatriz := make([]string, 3)

		if i, h := text.ContainsSliceSliceString(options.plugins, plugin[1]); !h {
			formOfMatriz[0] = plugin[1] // name
			formOfMatriz[1] = plugin[2] // version
			formOfMatriz[2] = plugin[0] // match

			options.plugins = append(options.plugins, formOfMatriz)
		} else {
			if options.plugins[i][1] == "" {
				options.plugins[i][1] = plugin[2]
			}
			if !strings.Contains(options.plugins[i][2], plugin[0]) {
				options.plugins[i][2] = options.plugins[i][2] + "ˆ" + plugin[0]
			}
		}
	}

	return options.plugins
}

// Aggressive :: Aggressive enumeration has 3 steps to complete.
// In the first she enumerates, in the second step she tries to identify the version using the passive mode and the third she tries to enumerate the version through brute-force in the theme files.
func (options *plugin) Aggressive() [][]string {
	if response := commons.DirectoryPlugins(); response.Response.StatusCode != 200 {
		regxp := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		for _, plugin := range regxp.FindAllStringSubmatch(response.Raw, -1) {
			formOfMatriz := make([]string, 3)

			if _, h := text.ContainsSliceSliceString(options.plugins, plugin[1]); !h {
				formOfMatriz[0] = plugin[1] // name

				options.plugins = append(options.plugins, formOfMatriz)
			}
		}
	}

	regxp := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/plugins/(.*?)/.*?[.css|.js]")

	for _, plugin := range regxp.FindAllStringSubmatch(options.raw, -1) {
		formOfMatriz := make([]string, 3)

		if _, h := text.ContainsSliceSliceString(options.plugins, plugin[1]); !h {
			formOfMatriz[0] = plugin[1] // name

			options.plugins = append(options.plugins, formOfMatriz)
		}
	}

	options.Passive()

	for _, pluginOfMatriz := range options.plugins {
		path := database.Memory.GetString("HTTP wp-content") + "/plugins/" + pluginOfMatriz[0] + "/"

		if match, version := extensions.GetVersionByIndexOf(path); version != "" {
			pluginOfMatriz[1] = version
			pluginOfMatriz[2] = pluginOfMatriz[2] + "ˆ" + match
		} else if match, version := extensions.GetVersionByReadme(path); version != "" {
			pluginOfMatriz[1] = version
			pluginOfMatriz[2] = pluginOfMatriz[2] + "ˆ" + match
		} else if match, version := extensions.GetVersionByChangeLogs(path); version != "" {
			pluginOfMatriz[1] = version
			pluginOfMatriz[2] = pluginOfMatriz[2] + "ˆ" + match
		} else if match, version := extensions.GetVersionByReleaseLog(path); version != "" {
			pluginOfMatriz[1] = version
			pluginOfMatriz[2] = pluginOfMatriz[2] + "ˆ" + match
		}
	}

	return options.plugins
}
