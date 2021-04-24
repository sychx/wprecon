package enumerate

import (
	"regexp"
	"sync"
	"time"

	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
)

const (
	MatchThemePassive            = "/themes/(.*?)/.*?[css|js].*?ver=(\\d{1,2}\\.\\d{1,2}\\.\\d{1,3})"
	MatchThemePassiveNoVersion   = "/themes/(.*?)/.*?[.css|.js]"
	MatchThemeAgressiveDirectory = "<a href=\"(.*?)/\">.*?/</a>"
)

type theme struct {
	target, raw, wpContentPath              string
	LenThemesPassive, LenThemesAggressive   int
}

func NewThemes(target, raw, wpcontent string) *theme {
	return &theme{target: target, raw: raw, wpContentPath: wpcontent}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (self *theme) Passive(channel chan [5]string) {
	var (
		mx sync.Mutex
		re          = regexp.MustCompile(self.wpContentPath + MatchThemePassive)
		allsubmatch = re.FindAllStringSubmatch(self.raw, -1)
		/*
		[0] submatch
		[1] name
		[2] version
		[3] type discovery
		[4] confidence
		*/
		themes     [][5]string
	)

	for _, submatch := range allsubmatch {
		mx.Lock()
		if index := findByValueInIndex(themes, submatch[1]); index == -1 {
			var form [5]string

			form[0] = submatch[0]
			form[1] = submatch[1]
			form[2] = submatch[2]
			form[3] = "Passive Enumerate"
			form[4] = "Enumerate Passive, in Index Page."

			themes = append(themes, form)
		} else {
			themes[index][0] = themes[index][0] + "," + submatch[0]
			themes[index][2] = themes[index][2] + "," + submatch[2]
		}
		mx.Unlock()
	}

	self.LenThemesPassive = len(themes)

	for _, theme := range themes {
		channel <- theme
	}
}

func (self *theme) Aggressive(channel chan [5]string) {
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
		themes [][5]string
	)

	wg.Add(2)
	
	go func() {
		var response = commons.DirectoryThemes()
		var re = regexp.MustCompile(MatchThemeAgressiveDirectory)

		for _, submatch := range re.FindAllStringSubmatch(response.Raw, -1) {
			/*
			If the condition is true, it means that the theme exists. Soon he will not add the list, but the condition will go to the else that will add another match for the themes.
			Note: For you to understand better, I recommend that you see this in operation.
			*/
			mx.Lock()
			if index, contains := findStringInSliceSlice(themes, 1, submatch[1]); !contains {
				var form [5]string
	
				form[0] = submatch[0]
				form[1] = submatch[1]
				form[3] = "Aggressive Enumerate"
				form[4] = "Enumerate By Index Of \"/wp-content/themes\""
					
				themes = append(themes, form)
			} else {
				themes[index][0] = themes[index][0] + "," + submatch[0]
			}
			mx.Unlock()
		}

		defer wg.Done()
	}()

	go func() {
		var channelx = make(chan [5]string)
		
		go self.Passive(channelx)

		time.Sleep(2*time.Second)

		for i := 1; i <= self.LenThemesPassive; i++ {
			select {
			case submatch := <- channelx:
				mx.Lock()
				if index := findByValueInIndex(themes, submatch[1]); index == -1 {
					themes = append(themes, submatch)
				} else {
					themes[index][0] = themes[index][0] + "," + submatch[0]
					themes[index][2] = themes[index][2] + "," + submatch[2]
				}
				mx.Unlock()
			}
		}

		defer wg.Done()
	}()

	wg.Wait()
	
	self.LenThemesPassive = 0
	self.LenThemesAggressive = len(themes)

	for _, theme := range themes {
		var path = self.wpContentPath + "/themes/" + theme[1] + "/"

		if match, version := extensions.GetVersionByIndexOf(self.target, path); version != "" {
			theme[0] = theme[0] + "," + match
			theme[2] = theme[2] + "," + version
			theme[3] = "Aggressive Enumerate"
		} else if match, version := extensions.GetVersionByReadme(self.target, path); version != "" {
			theme[0] = theme[0] + "," + match
			theme[2] = theme[2] + "," + version
			theme[3] = "Aggressive Enumerate"
		} else if match, version := extensions.GetVersionByChangeLogs(self.target, path); version != "" {
			theme[0] = theme[0] + "," + match
			theme[2] = theme[2] + "," + version
			theme[3] = "Aggressive Enumerate"
		} else if match, version := extensions.GetVersionByReleaseLog(self.target, path); version != "" {
			theme[0] = theme[0] + "," + match
			theme[2] = theme[2] + "," + version
			theme[3] = "Aggressive Enumerate"
 		}

		channel <- theme
	}
}
