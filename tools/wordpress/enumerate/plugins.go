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
	var p plugin

	p.raw = database.Memory.GetString("HTTP Index Raw")

	return &p
}

// Passive :: As the name says, this function will make an enumeration in an passive way.
// Passive enumeration may not be the best option when searching for vulnerabilities.
// The return is an matriz, containing 3 spaces 0 = name, 1 = version, 2 = match.
func (p *plugin) Passive() [][]string {
	rex := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/plugins/(.*?)/.*?[css|js].*?ver=([0-9\\.]*)")

	for _, plugin := range rex.FindAllStringSubmatch(p.raw, -1) {
		form := make([]string, 3)

		if i, h := text.ContainsSliceSliceString(p.plugins, plugin[1]); !h {
			form[0] = plugin[1] // name
			form[1] = plugin[2] // version
			form[2] = plugin[0] // match

			p.plugins = append(p.plugins, form)
		} else {
			if p.plugins[i][1] == "" {
				p.plugins[i][1] = plugin[2]
			}
			if !strings.Contains(p.plugins[i][2], plugin[0]) {
				p.plugins[i][2] = p.plugins[i][2] + "ˆ" + plugin[0]
			}
		}
	}

	return p.plugins
}

// PluginsEnumerateAgressive :: As the name says, this function will make an enumeration in an aggressive way.
// It will try to access the "wp-content/plugins" file if it does not have an index of, wprecon will use the PluginsEnumeratePassive function so that it can list the plugins.
// And when finished, it will send a list with the found plugins and their version.
// The return is an matriz, containing 3 spaces 0 = name, 1 = version, 2 = match.
func (p *plugin) Aggressive() [][]string {
	go func() {
		if response := commons.DirectoryPlugins(); response.Response.StatusCode != 200 {
			rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

			for _, plugin := range rex.FindAllStringSubmatch(response.Raw, -1) {
				form := make([]string, 3)

				if i, h := text.ContainsSliceSliceString(p.plugins, plugin[1]); !h {
					form[0] = plugin[1] // name
					form[2] = plugin[0] // match

					p.plugins = append(p.plugins, form)
				} else {
					if !strings.Contains(p.plugins[i][2], plugin[0]) {
						p.plugins[i][2] = p.plugins[i][2] + "ˆ" + plugin[0]
					}
				}
			}
		}
	}()

	rex := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/plugins/(.*?)/.*?[css|js]")

	for _, plugin := range rex.FindAllStringSubmatch(p.raw, -1) {
		form := make([]string, 3)

		if i, h := text.ContainsSliceSliceString(p.plugins, plugin[1]); !h {
			form[0] = plugin[1] // name
			form[2] = plugin[0] // match

			p.plugins = append(p.plugins, form)
		} else {
			if !strings.Contains(p.plugins[i][2], plugin[0]) {
				p.plugins[i][2] = p.plugins[i][2] + "ˆ" + plugin[0]
			}
		}
	}

	p.Passive()

	for _, ppp := range p.plugins {
		path := database.Memory.GetString("HTTP wp-content") + "/plugins/" + ppp[1] + "/"

		if match, version := extensions.GetVersionByIndexOf(path); version != "" {
			ppp[1] = version
			ppp[2] = match
		} else if match, version := extensions.GetVersionByReadme(path); version != "" {
			ppp[1] = version
			ppp[2] = match
		} else if match, version := extensions.GetVersionByChangeLogs(path); version != "" {
			ppp[1] = version
			ppp[2] = match
		} else if match, version := extensions.GetVersionByReleaseLog(path); version != "" {
			ppp[1] = version
			ppp[2] = match
		}
	}

	return p.plugins
}
