package internal

import (
	"encoding/json"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

type githubAPIJSON struct {
	Author struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Email       string `json:"email"`
	} `json:"Author"`
	App struct {
		Description     string `json:"description"`
		Version         string `json:"version"`
		VersionReselase string `json:"version-reselase"`
	} `json:"App"`
}

// UpdateCheck :: This function will be responsible for checking and printing on the screen whether there is an update or not.
func UpdateCheck() {
	var githubJSON githubAPIJSON

	printer.Loading("Checking Version!")

	request, _ := gohttp.HttpRequest(gohttp.Http{Method: "GET", URL: "https://raw.githubusercontent.com/blackcrw/wprecon/dev/internal/config.json"})

	if err := json.NewDecoder(request.Body).Decode(&githubJSON); err != nil {
		printer.LoadingDanger("An error occurred while trying to check the version.")
	} else if githubJSON.App.Version != Version() {
		printer.LoadingDone("There is a new version!", "New:", githubJSON.App.Version, "Download: https://github.com/blackcrw/wprecon/releases")
	}
}
