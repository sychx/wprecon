package extensions

import (
	"regexp"

	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
)

// CheckUpdate :: This function will be responsible for checking and printing on the screen whether there is an update or not.
func CheckUpdate() string {
	http := gohttp.NewHTTPClient().SetURLFull("https://raw.githubusercontent.com/blackbinn/wprecon/master/internal/pkg/extensions/version.go").SetSleep(0)

	var request, _ = http.Run()

	var re = regexp.MustCompile("var Version string = \"(.*?)\"")

	if findVersion := re.FindAllStringSubmatch(request.Raw, -1); findVersion[0][1] != Version {
		return findVersion[0][1]
	}

	return ""
}
