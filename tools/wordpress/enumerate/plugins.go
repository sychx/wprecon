package enumerate

import "regexp"

const (
	MatchPluginPassive            = "/plugins/(.*?)/.*?[css|js].*?ver=(\\d{1,2}\\.\\d{1,2}\\.\\d{1,3})"
	MatchPluginAgressiveDirectory = "<a href=\"(.*?)/\">.*?/</a>"
	MatchPluginAgressiveNoVersion = "/plugins/(.*?)/.*?[.css|.js]"
)

type plugin struct {
	raw           string
	wpContentPath string
}

func NewPlugins(raw string, wpcontent string) *plugin {
	return &plugin{raw: raw, wpContentPath: wpcontent}
}

// Passive :: "How does passive enumeration work?"
// We took the source code of the index that was saved in memory and from there we do a search using the regex.
func (object *plugin) Passive() [][]string {
	re := regexp.MustCompile(object.wpContentPath + MatchPluginPassive)

	return re.FindAllStringSubmatch(object.raw, -1)
}
