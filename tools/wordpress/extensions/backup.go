package extensions

import (
	"fmt"
	"regexp"
)

func FindBackupFileOrPath(raw string) (pathSlice []string) {
	rex := regexp.MustCompile("<a href=\"([back[wp|up|.*?]|bkp].*?)\">.*?</a>")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, plugin := range submatchall {
		path := fmt.Sprintf("%s", plugin[1])

		pathSlice = append(pathSlice, path)
	}

	return
}
