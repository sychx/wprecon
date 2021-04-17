package extensions

import (
	"regexp"
)

const (
	MatchFindBackupFileOrPatch = "<a href=\"([back[wp|up|.*?]|bkp].*?)\">.*?</a>"
)

// FindBackupFileOrPath ::
// Revisar o funcionamento !!!
func FindBackupFileOrPath(raw string) []string {
	var slice []string

	var rex = regexp.MustCompile(MatchFindBackupFileOrPatch)

	for _, plugin := range rex.FindAllStringSubmatch(raw, -1) {
		slice = append(slice, plugin[1])
	}

	return slice
}
