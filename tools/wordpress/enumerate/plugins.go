package enumerate

import (
	"regexp"
	"sync"
	"time"

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
func (self *plugin) Passive(channel chan [5]string) {
	var (
		mx sync.Mutex
		re          = regexp.MustCompile(self.wpContentPath + MatchPluginPassive)
		allsubmatch = re.FindAllStringSubmatch(self.raw, -1)
		/*
		[0] submatch
		[1] name
		[2] version
		[3] type discovery
		[4] confidence
		*/
		plugins     [][5]string
	)

	for _, submatch := range allsubmatch {
		mx.Lock()
		if index := findByValueInIndex(plugins, submatch[1]); index == -1 {
			var form [5]string

			form[0] = submatch[0]
			form[1] = submatch[1]
			form[2] = submatch[2]
			form[3] = "Passive Enumerate"
			form[4] = "Enumerate Passive, in Index Page."

			plugins = append(plugins, form)
		} else {
			plugins[index][0] = plugins[index][0] + "," + submatch[0]
			plugins[index][2] = plugins[index][2] + "," + submatch[2]
		}
		mx.Unlock()
	}

	self.LenPluginsPassive = len(plugins)

	for _, plugin := range plugins {
		channel <- plugin
	}
}

func (self *plugin) Aggressive(channel chan [5]string) {
	var (
		wg sync.WaitGroup
		mx sync.Mutex
		/*
		[0] submatch
		[1] name
		[2] version
		[3] type discovery
		[4] confidence
		*/
		plugins [][5]string
	)

	wg.Add(2)
	
	go func() {
		var response = commons.DirectoryPlugins()
		var re = regexp.MustCompile(MatchPluginAgressiveDirectory)

		for _, submatch := range re.FindAllStringSubmatch(response.Raw, -1) {
			/*
			If the condition is true, it means that the plugin exists. Soon he will not add the list, but the condition will go to the else that will add another match for the plugins.
			Note: For you to understand better, I recommend that you see this in operation.
			*/
			mx.Lock()
			if index, contains := findStringInSliceSlice(plugins, 1, submatch[1]); !contains {
				var form [5]string
	
				form[0] = submatch[0]
				form[1] = submatch[1]
				form[3] = "Aggressive Enumerate"
				form[4] = "Enumerate By Index Of \"/wp-content/plugins\""
					
				plugins = append(plugins, form)
			} else {
				plugins[index][0] = plugins[index][0] + "," + submatch[0]
			}
			mx.Unlock()
		}

		defer wg.Done()
	}()

	go func() {
		var channelx = make(chan [5]string)
		
		go self.Passive(channelx)

		time.Sleep(2*time.Second)

		for i := 1; i <= self.LenPluginsPassive; i++ {
			select {
			case submatch := <- channelx:
				mx.Lock()
				if index := findByValueInIndex(plugins, submatch[1]); index == -1 {
					plugins = append(plugins, submatch)
				} else {
					plugins[index][0] = plugins[index][0] + "," + submatch[0]
					plugins[index][2] = plugins[index][2] + "," + submatch[2]
				}
				mx.Unlock()
			}
		}

		defer wg.Done()
	}()

	wg.Wait()
	
	self.LenPluginsPassive = 0
	self.LenPluginsAggressive = len(plugins)

	for _, plugin := range plugins {
		var path = self.wpContentPath + "/plugins/" + plugin[1] + "/"

		if match, version := extensions.GetVersionByIndexOf(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
			plugin[3] = "Aggressive Enumerate"
		} else if match, version := extensions.GetVersionByReadme(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
			plugin[3] = "Aggressive Enumerate"
		} else if match, version := extensions.GetVersionByChangeLogs(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
			plugin[3] = "Aggressive Enumerate"
		} else if match, version := extensions.GetVersionByReleaseLog(self.target, path); version != "" {
			plugin[0] = plugin[0] + "," + match
			plugin[2] = plugin[2] + "," + version
			plugin[3] = "Aggressive Enumerate"
 		}

		channel <- plugin
	}
}
