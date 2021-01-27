package update

import (
	"encoding/json"

	"github.com/blackbinn/wprecon/internal/pkg/version"
	"github.com/blackbinn/wprecon/pkg/gohttp"
)

type githubAPIJSON struct {
	App struct {
		Description string `json:"description"`
		Version     string `json:"version"`
		Endpoint    string `json:"endpoint"`
	} `json:"App"`
}

// CheckUpdate :: This function will be responsible for checking and printing on the screen whether there is an update or not.
func CheckUpdate() string {
	var githubJSON githubAPIJSON

	http := gohttp.NewHTTPClient().SetURLFull("https://raw.githubusercontent.com/blackbinn/wprecon/dev/internal/config/config.json")

	request, _ := http.Run()

	json.Unmarshal([]byte(request.Raw), &githubJSON)

	if githubJSON.App.Version != version.Version {
		return githubJSON.App.Version
	}

	return ""
}
