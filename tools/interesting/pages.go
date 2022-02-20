package interesting

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

func AdminPage() (models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory("wp-admin/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return models.InterestingModel{}, err }
	
	var model = models.InterestingModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}
	
	defer database.Memory.SetString("HTTP Admin Page Status", response.Response.Status)
	
	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model, nil
}

func RobotsPage() (models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory("robots.txt")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return models.InterestingModel{}, err }

	var model = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model, nil
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.sitemap.xml.status"]
func SitemapPage() (models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory("sitemap.xml")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { printer.Danger(fmt.Sprintf("%s", err)) }

	var model = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model, nil
}

func ReadmePage() (models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory("readme.html")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return models.InterestingModel{}, err }

	var model = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model, nil
}
