package enumerate

import (
	"fmt"
	"regexp"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
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
	if response := commons.DirectoryPlugins(); response.Raw != "" {
		rex := regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		submatchall := rex.FindAllSubmatch([]byte(response.Raw), -1)

		for _, plugin := range submatchall {
			name := fmt.Sprintf("%s", plugin[1])

			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = ""
		}
	} else if pluginslist := PluginsEnumeratePassive(); len(pluginslist) > 0 {
	} else if len(pluginslist) == 0 {
		raw := InfosWprecon.OtherInformationsString["target.http.index.raw"]

		rex := regexp.MustCompile("/wp-content/plugins/(.*?)/.*?[css|js]")
		submatchall := rex.FindAllSubmatch([]byte(raw), -1)

		for _, plugin := range submatchall {
			name := fmt.Sprintf("%s", plugin[1])

			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = ""
		}
	} else {
		return make(map[string]string)
	}

	for name := range InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"] {
		path := "/wp-content/plugins/" + name + "/"

		if version := extensions.GetVersionByIndexOf(path); version != "" {
			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
		} else if version := extensions.GetVersionByReadme(path); version != "" {
			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
		} else if version := extensions.GetVersionByChangeLogs(path); version != "" {
			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
		} else if version := extensions.GetVersionByReleaseLog(path); version != "" {
			InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"][name] = version
		}
	}

	return InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"]
}
