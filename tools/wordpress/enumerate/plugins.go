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
	target, raw, wpContentPath              string
	LenPluginsPassive, LenPluginsAggressive int
}

func NewPlugins(target, raw, wpcontent string) *plugin {
	return &plugin{target: target, raw: raw, wpContentPath: wpcontent}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (self *plugin) Passive() [][]string {
	var (
		re          = regexp.MustCompile(self.wpContentPath + MatchPluginPassive)
		allsubmatch = re.FindAllStringSubmatch(self.raw, -1)
		plugins     [][]string
	)

	self.LenPluginsPassive = len(allsubmatch)

	for _, submatch := range allsubmatch {
		if _, has := text.ContainsInSliceSlice(plugins, submatch[1]); !has {
			plugins = append(plugins, submatch)
		} else if index := text.FindByValueInIndex(plugins, submatch[1]); index != -1 {
			plugins[index][0] = plugins[index][0] + "," + submatch[0]
			plugins[index][2] = plugins[index][2] + "," + submatch[2]
		}
	}

	return plugins
}

func (self *plugin) Aggressive(channel chan []string) {
	var (
		wg      sync.WaitGroup
		mx      sync.Mutex
		plugins [][]string
	)

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
				if _, contains := text.FindStringInSliceSlice(plugins, 1, submatch[1]); !contains {
					plugins = append(plugins, submatch)
				} else if index := text.FindByValueInIndex(plugins, submatch[1]); index != -1 {
					plugins[index][0] = plugins[index][0] + "," + submatch[0]
				}
				mx.Unlock()
			}
		}

		defer wg.Done()
	}()

	go func() {
		for _, submatch := range self.Passive() {
			mx.Lock()
			if _, has := text.ContainsInSliceSlice(plugins, submatch[1]); !has {
				plugins = append(plugins, submatch)
			} else if index := text.FindByValueInIndex(plugins, submatch[1]); index != -1 {
				plugins[index][0] = plugins[index][0] + "," + submatch[0]
				plugins[index][2] = plugins[index][2] + "," + submatch[2]
			}
			mx.Unlock()
		}

		defer wg.Done()
	}()

	wg.Wait()

	self.LenPluginsAggressive = len(plugins)

	for _, plugin := range plugins {
		var path = self.wpContentPath + "/plugins/" + plugin[1] + "/"

		if match, version := extensions.GetVersionByIndexOf(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		} else if match, version := extensions.GetVersionByReadme(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		} else if match, version := extensions.GetVersionByChangeLogs(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		} else if match, version := extensions.GetVersionByReleaseLog(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
		}

		channel <- []string{plugin[0], plugin[1], plugin[2]}
	}

	close(channel)
}
