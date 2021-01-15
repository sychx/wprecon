package enumerate

import (
	"fmt"
	"regexp"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/wordlist"
	"github.com/blackcrw/wprecon/tools/wordpress/commons"
	"github.com/blackcrw/wprecon/tools/wordpress/extensions"
)

// PluginsEnumeratePassive :: As the name says, this function will make an enumeration in an passive way.
// Passive enumeration may not be the best option when searching for vulnerabilities.
// (I don't recommend) 40% confidence.
func PluginsEnumeratePassive() map[string]string {
	raw := InfosWprecon.OtherInformationsString["target.http.index.raw"]

	rex := regexp.MustCompile("/wp-content/plugins/(.*?)/.*?[css|js].*?ver=([0-9\\.]*)")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, plugin := range submatchall {
		name := fmt.Sprintf("%s", plugin[1])
		version := fmt.Sprintf("%s", plugin[2])

		InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
	}

	return InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"]
}

// PluginsEnumerateAgressive :: As the name says, this function will make an enumeration in an aggressive way.
// It will try to access the "wp-content/plugins" file if it does not have an index of, wprecon will use the PluginsEnumeratePassive function so that it can list the plugins.
// And when finished, it will send a list with the found plugins and their version.
// The plugins will be returned based on this list: InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"]
// (Recommend) 80% confidence
func PluginsEnumerateAgressive() map[string]string {
	commons.DirectoryPlugins()
	PluginsEnumeratePassive()

	if InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"] != "" && len(InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"]) <= 4 {
		rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		submatchall := rex.FindAllSubmatch([]byte(InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]), -1)

		for _, plugin := range submatchall {
			name := fmt.Sprintf("%s", plugin[1])

			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = "Not found"
		}
	}

	for name := range InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"] {
		done := false

		if InfosWprecon.Verbose == true {
			go func() {
				response := extensions.SimpleRequest(InfosWprecon.Target, fmt.Sprintf("wp-content/plugins/%s", name))
				extensions.GetFileExtensions(response.URL.Full, response.Raw)
			}()
		}

		for _, value := range wordlist.WPchangesLogs {
			dir := fmt.Sprintf("/wp-content/plugins/%s/%s", name, value)

			if response := extensions.SimpleRequest(InfosWprecon.Target, dir); response.StatusCode == 200 && response.Raw != "" {
				if version := extensions.GetVersionChangelog(response.Raw); version != "" {
					InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
					done = true
					break
				}
			}
		}

		if done == false {
			for _, value := range wordlist.WPreadme {
				dir := fmt.Sprintf("wp-content/plugins/%s/%s", name, value)

				if response := extensions.SimpleRequest(InfosWprecon.Target, dir); response.StatusCode == 200 && response.Raw != "" {
					if version := extensions.GetVersionStableTag(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
						done = true
						break
					} else if version := extensions.GetVersionChangelog(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
						done = true
						break
					}
				}
			}
		}

		if done == false {
			for _, value := range wordlist.WPreleaseLog {
				dir := fmt.Sprintf("wp-content/plugins/%s/%s", name, value)

				if response := extensions.SimpleRequest(InfosWprecon.Target, dir); response.StatusCode == 200 && response.Raw != "" {
					if version := extensions.GetVersionStableTag(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
						done = true
						break
					} else if version := extensions.GetVersionChangelog(response.Raw); version != "" {
						InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
						done = true
						break
					}
				}
			}
		}
	}

	return InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"]
}
