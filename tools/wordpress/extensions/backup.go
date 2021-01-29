package extensions

import (
	"fmt"
	"regexp"
)

func FindBackupFileOrPath(raw string) []string {
	var pathSlice []string

	rex := regexp.MustCompile("<a href=\"(back[wp|up].*?[backup|.*?].*?)\">.*?</a>")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, plugin := range submatchall {
		path := fmt.Sprintf("%s", plugin[1])

		pathSlice = append(pathSlice, path)
	}

	return pathSlice
}
