package enumerate

import (
	"fmt"
	"regexp"

	. "github.com/blackbinn/wprecon/cli/config"
)

func WordpressVersionPassive() string {
	raw := InfosWprecon.OtherInformationsString["target.http.index.raw"]

	rex := regexp.MustCompile("<meta name=\"generator\" content=\"WordPress ([0-9.-]*).*?")

	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	for _, slicebytes := range submatchall {
		version := fmt.Sprintf("%s", slicebytes[1])

		InfosWprecon.OtherInformationsString["target.http.wordpress.version"] = version
	}

	return InfosWprecon.OtherInformationsString["target.http.wordpress.version"]
}

func WordpressVersionAggressive() {

}
