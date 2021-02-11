package enumerate

import (
	"fmt"
	"regexp"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
)

// ThemesEnumeratePassive :: As the name says, this function will make an enumeration in an passive way.
// Passive enumeration may not be the best option when searching for vulnerabilities.
// (I don't recommend) 40% confidence.
func ThemesEnumeratePassive() map[string]string {
	raw := Database.OtherInformationsString["target.http.index.raw"]

	rex := regexp.MustCompile(Database.WPContent + "/themes/(.*?)/.*?[css|js].*?ver=([0-9\\.]*)")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, theme := range submatchall {
		name := fmt.Sprintf("%s", theme[1])
		version := fmt.Sprintf("%s", theme[2])

		Database.OtherInformationsMapString["target.http.themes.versions"][name] = version
	}

	return Database.OtherInformationsMapString["target.http.themes.versions"]
}

// ThemesEnumerateAgressive :: As the name says, this function will make an enumeration in an aggressive way.
// It will try to access the "wp-content/themes" file if it does not have an index of, wprecon will use the ThemesEnumeratePassive function so that it can list the themes.
// And when finished, it will send a list with the found themes and their version.
// The themes will be returned based on this list: Database.OtherInformationsMapString["target.http.themes.versions"]
func ThemesEnumerateAgressive() map[string]string {
	if response := commons.DirectoryThemes(); response.Response.StatusCode == 200 && response.Raw != "" {
		rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		submatchall := rex.FindAllSubmatch([]byte(Database.OtherInformationsString["target.http.wp-content/themes.indexof.raw"]), -1)

		for _, theme := range submatchall {
			name := fmt.Sprintf("%s", theme[1])

			Database.OtherInformationsMapString["target.http.themes.versions"][name] = ""
		}

		ThemesEnumeratePassive()
	} else if themesList := ThemesEnumeratePassive(); len(themesList) > 0 {
	} else if len(themesList) == 0 {
		raw := Database.OtherInformationsString["target.http.index.raw"]

		rex := regexp.MustCompile(Database.WPContent + "/themes/(.*?)/.*?[css|js]")
		submatchall := rex.FindAllSubmatch([]byte(raw), -1)

		for _, theme := range submatchall {
			name := fmt.Sprintf("%s", theme[1])

			Database.OtherInformationsMapString["target.http.themes.versions"][name] = ""
		}
	} else {
		return make(map[string]string)
	}

	for name := range Database.OtherInformationsMapString["target.http.themes.versions"] {
		path := Database.WPContent + "/themes/" + name + "/"

		if version := extensions.GetVersionByIndexOf(path); version != "" {
			Database.OtherInformationsMapString["target.http.themes.versions"][name] = version
		} else if version := extensions.GetVersionByReadme(path); version != "" {
			Database.OtherInformationsMapString["target.http.themes.versions"][name] = version
		} else if version := extensions.GetVersionByChangeLogs(path); version != "" {
			Database.OtherInformationsMapString["target.http.themes.versions"][name] = version
		} else if version := extensions.GetVersionByReleaseLog(path); version != "" {
			Database.OtherInformationsMapString["target.http.themes.versions"][name] = version
		}
	}

	return Database.OtherInformationsMapString["target.http.themes.versions"]
}
