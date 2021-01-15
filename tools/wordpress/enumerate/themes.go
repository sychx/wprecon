package enumerate

import (
	"fmt"
	"regexp"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/wordlist"
	"github.com/blackcrw/wprecon/tools/wordpress/commons"
	"github.com/blackcrw/wprecon/tools/wordpress/extensions"
)

// ThemesEnumeratePassive :: As the name says, this function will make an enumeration in an passive way.
// Passive enumeration may not be the best option when searching for vulnerabilities.
// (I don't recommend) 40% confidence.
func ThemesEnumeratePassive() map[string]string {
	raw := InfosWprecon.OtherInformationsString["target.http.index.raw"]

	rex := regexp.MustCompile("/wp-content/themes/(.*?)/.*?[css|js].*?ver=([0-9\\.]*)")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, plugin := range submatchall {
		name := fmt.Sprintf("%s", plugin[1])
		version := fmt.Sprintf("%s", plugin[2])

		InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][name] = version
	}

	return InfosWprecon.OtherInformationsMapString["target.http.themes.versions"]
}

// ThemesEnumerateAgressive :: As the name says, this function will make an enumeration in an aggressive way.
// It will try to access the "wp-content/themes" file if it does not have an index of, wprecon will use the ThemesEnumeratePassive function so that it can list the themes.
// And when finished, it will send a list with the found themes and their version.
// The themes will be returned based on this list: InfosWprecon.OtherInformationsMapString["target.http.themes.versions"]
func ThemesEnumerateAgressive() map[string]string {
	commons.DirectoryThemes()
	ThemesEnumeratePassive()

	if InfosWprecon.OtherInformationsString["target.http.wp-content/themes.indexof.raw"] != "" && len(InfosWprecon.OtherInformationsMapString["target.http.themes.versions"]) <= 0 {
		rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		submatchall := rex.FindAllSubmatch([]byte(InfosWprecon.OtherInformationsString["target.http.wp-content/themes.indexof.raw"]), -1)

		for _, plugin := range submatchall {
			name := fmt.Sprintf("%s", plugin[1])

			InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][name] = "0"
		}
	}

	for key := range InfosWprecon.OtherInformationsMapString["target.http.themes.versions"] {
		done := false

		for _, value := range wordlist.WPchangesLogs {
			dir := fmt.Sprintf("/wp-content/themes/%s/%s", key, value)

			if response := extensions.SimpleRequest(InfosWprecon.Target, dir); response.StatusCode == 200 && response.Raw != "" {
				if version := extensions.GetVersionChangelog(response.Raw); version != "" {
					InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][key] = version
					done = true
					break
				}
			}
		}

		if done == false {
			for _, value := range wordlist.WPreadme {
				dir := fmt.Sprintf("/wp-content/themes/%s/%s", key, value)

				if response := extensions.SimpleRequest(InfosWprecon.Target, dir); response.StatusCode == 200 && response.Raw != "" {

					if version := extensions.GetVersionStableTag(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][key] = version
						done = true
						break
					} else if version := extensions.GetVersionChangelog(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][key] = version
						done = true
						break
					}
				}
			}
		}

		if done == false {
			for _, value := range wordlist.WPreleaseLog {
				dir := fmt.Sprintf("/wp-content/themes/%s/%s", key, value)

				if response := extensions.SimpleRequest(InfosWprecon.Target, dir); response.StatusCode == 200 && response.Raw != "" {

					if version := extensions.GetVersionStableTag(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][key] = version
						done = true
						break
					} else if version := extensions.GetVersionChangelog(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.themes.versions"][key] = version
						done = true
						break
					}
				}
			}
		}
	}

	return InfosWprecon.OtherInformationsMapString["target.http.themes.versions"]
}
