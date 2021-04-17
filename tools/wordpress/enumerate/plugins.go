package enumerate

import (
	"regexp"
	"sync"

	"github.com/blackbinn/wprecon/internal/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
)

const (
	MatchPluginPassive            = "/plugins/(.*?)/.*?[css|js].*?ver=(\\d{1,2}\\.\\d{1,2}\\.\\d{1,3})"
	MatchPluginPassiveNoVersion   = "/plugins/(.*?)/.*?[.css|.js]"
	MatchPluginAgressiveDirectory = "<a href=\"(.*?)/\">.*?/</a>"
)

type plugin struct {
	target, raw, wpContentPath                  string
	plugins                                     [][]string
	CountPluginsPassive, CountPluginsAggressive int
}

func NewPlugins(target, raw, wpcontent string) *plugin {
	return &plugin{target: target, raw: raw, wpContentPath: wpcontent}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (object *plugin) Passive() [][]string {
	var re = regexp.MustCompile(object.wpContentPath + MatchPluginPassive)
	var submatch = re.FindAllStringSubmatch(object.raw, -1)

	object.CountPluginsPassive = len(submatch)

	return submatch
}

func (object *plugin) Aggressive() [][]string {
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(2)

	go func() {
		if response := commons.DirectoryPlugins(); response.Response.StatusCode == 200 {
			var re = regexp.MustCompile(MatchPluginAgressiveDirectory)

			for _, submatch := range re.FindAllStringSubmatch(response.Raw, -1) {
				/*
					If the condition is true, it means that the plugin exists. Soon he will not add the list, but the condition will go to the else that will add another match for the plugins.
					Note: For you to understand better, I recommend that you see this in operation.
				*/
				mx.Lock()
				if _, contains := text.FindStringInSliceSlice(object.plugins, 1, submatch[1]); !contains {
					object.plugins = append(object.plugins, submatch)
				} else if index := text.FindByValueInIndex(object.plugins, submatch[1]); index != -1 {
					object.plugins[index][0] = object.plugins[index][0] + "," + submatch[0]
				}
				mx.Unlock()
			}
		}

		defer wg.Done()
	}()

	go func() {
		for _, submatch := range object.Passive() {
			mx.Lock()
			if _, has := text.ContainsInSliceSlice(object.plugins, submatch[1]); !has {
				object.plugins = append(object.plugins, submatch)
			} else if index := text.FindByValueInIndex(object.plugins, submatch[1]); index != -1 {
				object.plugins[index][0] = object.plugins[index][0] + "," + submatch[0]
				object.plugins[index][2] = object.plugins[index][2] + "," + submatch[2]
			}
			mx.Unlock()
		}

		defer wg.Done()
	}()

	wg.Wait()

	for _, plugin := range object.plugins {
		var path = object.wpContentPath + "/plugins/" + plugin[1] + "/"

		if match, version := extensions.GetVersionByIndexOf(object.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		} else if match, version := extensions.GetVersionByReadme(object.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		} else if match, version := extensions.GetVersionByChangeLogs(object.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		} else if match, version := extensions.GetVersionByReleaseLog(object.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		}
	}

	object.CountPluginsAggressive = len(object.plugins)

	return object.plugins
}
