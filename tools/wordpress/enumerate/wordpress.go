package enumerate

import (
	"fmt"
	"regexp"

	"github.com/blackbinn/wprecon/internal/database"
)

func WordpressVersionPassive() string {
	raw := database.Memory.GetString("HTTP Index Raw")

	rex := regexp.MustCompile("<meta name=\"generator\" content=\"WordPress ([0-9.-]*).*?")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, slicebytes := range submatchall {
		version := fmt.Sprintf("%s", slicebytes[1])

		database.Memory.SetString("HTTP WordPress Version", version)
	}

	return database.Memory.GetString("HTTP WordPress Version")
}
