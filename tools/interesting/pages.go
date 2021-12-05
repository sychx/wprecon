package interesting

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

func AdminPage() models.InterestingModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("wp-admin/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}
	
	var model = models.InterestingModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}
	
	defer database.Memory.SetString("HTTP Admin Page Status", response.Response.Status)
	
	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

func RobotsPage() models.InterestingModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("robots.txt")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	var model = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.sitemap.xml.status"]
func SitemapPage() models.InterestingModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("sitemap.xml")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	var model = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

func ReadmePage() models.InterestingModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("readme.html")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	var model = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}
