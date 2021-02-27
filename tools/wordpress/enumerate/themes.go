package enumerate

import (
	"regexp"
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
)

type theme struct {
	raw    string
	themes [][]string // matriz
}

func NewThemes() *theme {
	var p theme

	p.raw = database.Memory.GetString("HTTP Index Raw")

	return &p
}

// Passive :: As the name says, this function will make an enumeration in an passive way.
// Passive enumeration may not be the best option when searching for vulnerabilities.
// (I don't recommend) 40% confidence.
func (p *theme) Passive() [][]string {
	rex := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/themes/(.*?)/.*?[css|js].*?ver=([0-9\\.]*)")

	for _, theme := range rex.FindAllStringSubmatch(p.raw, -1) {
		form := make([]string, 3)

		if i, h := text.ContainsSliceSliceString(p.themes, theme[1]); !h {
			form[0] = theme[1] // name
			form[1] = theme[2] // version
			form[2] = theme[0] // match

			p.themes = append(p.themes, form)
		} else {
			if p.themes[i][1] == "" {
				p.themes[i][1] = theme[2]
			}
			if !strings.Contains(p.themes[i][2], theme[0]) {
				p.themes[i][2] = p.themes[i][2] + "ˆ" + theme[0]
			}
		}
	}

	return p.themes
}

// Aggressive :: As the name says, this function will make an enumeration in an aggressive way.
// It will try to access the "wp-content/themes" file if it does not have an index of, wprecon will use the ThemesEnumeratePassive function so that it can list the themes.
// And when finished, it will send a list with the found themes and their version.
func (p *theme) Aggressive() [][]string {
	go func() {
		if response := commons.DirectoryPlugins(); response.Response.StatusCode != 200 {
			rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

			for _, theme := range rex.FindAllStringSubmatch(response.Raw, -1) {
				form := make([]string, 3)

				if i, h := text.ContainsSliceSliceString(p.themes, theme[1]); !h {
					form[0] = theme[1] // name
					form[2] = theme[0] // match

					p.themes = append(p.themes, form)
				} else {
					if !strings.Contains(p.themes[i][2], theme[0]) {
						p.themes[i][2] = p.themes[i][2] + "ˆ" + theme[0]
					}
				}
			}
		}
	}()

	rex := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/themes/(.*?)/.*?[css|js]")

	for _, theme := range rex.FindAllStringSubmatch(p.raw, -1) {
		form := make([]string, 3)

		if i, h := text.ContainsSliceSliceString(p.themes, theme[1]); !h {
			form[0] = theme[1] // name
			form[2] = theme[0] // match

			p.themes = append(p.themes, form)
		} else {
			if !strings.Contains(p.themes[i][2], theme[0]) {
				p.themes[i][2] = p.themes[i][2] + "ˆ" + theme[0]
			}
		}
	}

	p.Passive()

	for _, ppp := range p.themes {
		path := database.Memory.GetString("HTTP wp-content") + "/themes/" + ppp[1] + "/"

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

	return p.themes
}
