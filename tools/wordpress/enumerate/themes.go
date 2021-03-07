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
	raw    string     // Index Raw Code
	themes [][]string // Matriz
}

func NewThemes() *theme {
	return &theme{raw: database.Memory.GetString("HTTP Index Raw")}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (options *theme) Passive() [][]string {
	regxp := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/themes/(.*?)/.*?[css|js]?ver=(\\d{1,2}\\.\\d{1,2}\\.\\d{1,3})")

	for _, theme := range regxp.FindAllStringSubmatch(options.raw, -1) {
		formOfMatriz := make([]string, 3)

		if count, exists := text.ContainsSliceSliceString(options.themes, theme[1]); !exists {
			formOfMatriz[0] = theme[1] // name
			formOfMatriz[1] = theme[2] // version
			formOfMatriz[2] = theme[0] // match

			options.themes = append(options.themes, formOfMatriz)
		} else {
			if options.themes[count][1] == "" {
				options.themes[count][1] = theme[2]
			}
			if !strings.Contains(options.themes[count][2], theme[0]) {
				options.themes[count][2] = options.themes[count][2] + "ˆ" + theme[0]
			}
		}
	}

	return options.themes
}

// Aggressive :: Aggressive enumeration has 3 steps to complete.
// In the first she enumerates, in the second step she tries to identify the version using the passive mode and the third she tries to enumerate the version through brute-force in the theme files.
func (options *theme) Aggressive() [][]string {
	go func() {
		if response := commons.DirectoryPlugins(); response.Response.StatusCode != 200 {
			regxp := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

			for _, theme := range regxp.FindAllStringSubmatch(response.Raw, -1) {
				formOfMatriz := make([]string, 3)

				if _, h := text.ContainsSliceSliceString(options.themes, theme[1]); !h {
					formOfMatriz[0] = theme[1] // name

					options.themes = append(options.themes, formOfMatriz)
				}
			}
		}
	}()

	regxp := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/themes/(.*?)/.*?[css|js]")

	for _, theme := range regxp.FindAllStringSubmatch(options.raw, -1) {
		formOfMatriz := make([]string, 3)

		if _, h := text.ContainsSliceSliceString(options.themes, theme[1]); !h {
			formOfMatriz[0] = theme[1] // name

			options.themes = append(options.themes, formOfMatriz)
		}
	}

	options.Passive()

	for _, themeOfMatriz := range options.themes {
		path := database.Memory.GetString("HTTP wp-content") + "/themes/" + themeOfMatriz[0] + "/"

		if match, version := extensions.GetVersionByIndexOf(path); version != "" {
			themeOfMatriz[1] = version
			themeOfMatriz[2] = match
		} else if match, version := extensions.GetVersionByReadme(path); version != "" {
			themeOfMatriz[1] = version
			themeOfMatriz[2] = match
		} else if match, version := extensions.GetVersionByChangeLogs(path); version != "" {
			themeOfMatriz[1] = version
			themeOfMatriz[2] = match
		} else if match, version := extensions.GetVersionByReleaseLog(path); version != "" {
			themeOfMatriz[1] = version
			themeOfMatriz[2] = match
		}
	}

	return options.themes
}
