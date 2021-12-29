package update

import (
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"gopkg.in/yaml.v2"
)

// Check :: This function will be responsible for checking and printing on the screen whether there is an update or not.
func GetVersion() string {
	var model models.ConfigModel

	var request = net.NewNETClient()
	request.SetURLFull("https://raw.githubusercontent.com/blackcrw/wprecon/wprecon-v2/internal/config/config.yaml")
	request.OnRandomUserAgent(true)

	var response, _ = request.Runner()

	yaml.Unmarshal([]byte(response.Raw), &model)

	return model.App.Version
}
