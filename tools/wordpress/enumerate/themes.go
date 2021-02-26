package enumerate

import (
	"regexp"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
)

// ThemesEnumeratePassive :: As the name says, this function will make an enumeration in an passive way.
// Passive enumeration may not be the best option when searching for vulnerabilities.
// (I don't recommend) 40% confidence.
func ThemesEnumeratePassive() map[string]string {
	raw := database.Memory.GetString("HTTP Index Raw")

	rex := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/themes/(.*?)/.*?[css|js].*?ver=([0-9\\.]*)")

	for _, theme := range rex.FindAllStringSubmatch(raw, -1) {
		database.Memory.SetMapMapString("HTTP Themes Versions", theme[1], theme[2])
	}

	return database.Memory.GetMapString("HTTP Themes Versions")
}

// ThemesEnumerateAgressive :: As the name says, this function will make an enumeration in an aggressive way.
// It will try to access the "wp-content/themes" file if it does not have an index of, wprecon will use the ThemesEnumeratePassive function so that it can list the themes.
// And when finished, it will send a list with the found themes and their version.
// The themes will be returned based on this list: Database.OtherInformationsMapString["target.http.themes.versions"]
func ThemesEnumerateAgressive() map[string]string {
	if response := commons.DirectoryThemes(); response.Response.StatusCode == 200 {
		rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		submatchall := rex.FindAllStringSubmatch(response.Raw, -1)

		for _, theme := range submatchall {
			database.Memory.SetMapMapString("HTTP Themes Versions", theme[1], "")
		}

		ThemesEnumeratePassive()
	} else if themesList := ThemesEnumeratePassive(); len(themesList) > 0 {
	} else if len(themesList) == 0 {
		raw := database.Memory.GetString("HTTP Index Raw")

		rex := regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/themes/(.*?)/.*?[css|js]")

		for _, theme := range rex.FindAllStringSubmatch(raw, -1) {
			database.Memory.SetMapMapString("HTTP Themes Versions", theme[1], "")
		}
	} else {
		return make(map[string]string)
	}

	for name := range database.Memory.GetMapString("HTTP Themes Versions") {
		path := database.Memory.GetString("HTTP wp-content") + "/themes/" + name + "/"

		if _, version := extensions.GetVersionByIndexOf(path); version != "" {
			database.Memory.SetMapMapString("HTTP Themes Versions", name, version)
		} else if _, version := extensions.GetVersionByReadme(path); version != "" {
			database.Memory.SetMapMapString("HTTP Themes Versions", name, version)
		} else if _, version := extensions.GetVersionByChangeLogs(path); version != "" {
			database.Memory.SetMapMapString("HTTP Themes Versions", name, version)
		} else if _, version := extensions.GetVersionByReleaseLog(path); version != "" {
			database.Memory.SetMapMapString("HTTP Themes Versions", name, version)
		}
	}

	return database.Memory.GetMapString("HTTP Themes Versions")
}
