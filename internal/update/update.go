package update

import (
	"encoding/json"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
)

// Check :: This function will be responsible for checking and printing on the screen whether there is an update or not.
func Check() string {
	var model models.VersionApiModel

	var response = net.SimpleRequest("https://raw.githubusercontent.com/blackbinn/wprecon/master/internal/config/config.json")

	json.Unmarshal([]byte(response.Raw), &model)

	if model.Configure.Version != "1" {
		return model.Configure.Version
	}

	return ""
}
