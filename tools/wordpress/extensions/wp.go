package extensions

import (
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
)

// WordpressCheck :: This function should be used to perform wordpress detection.
// "How does this detection work?", I decided to make a 'percentage system' where I will check if each item in a list exists... and if that item exists it will add +1 to accuracy.
// With "16.6" hits he says that wordpress is already detected. But it opens up an opportunity for you to choose whether to continue or not, because you are not 100% sure.
func WordpressCheck() float32 {
	var confidence float32
	var payloads = [4]string{
		"<meta name=\"generator content=\"WordPress",
		"<a href=\"http://www.wordpress.com\">Powered by WordPress</a>",
		"<link rel=\"https://api.wordpress.org/",
		"<link rel=\"https://api.w.org/\""}

	if has, _ := commons.AdminPage(); has == "true" || has == "redirect" {
		confidence++
	}
	if response := commons.DirectoryPlugins(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}
	if response := commons.DirectoryThemes(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}
	if response := commons.DirectoryUploads(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}

	for _, payload := range payloads {
		if strings.Contains(database.Memory.GetString("HTTP Index Raw"), payload) {
			confidence++
		}
	}

	return confidence / 8 * 100
}
