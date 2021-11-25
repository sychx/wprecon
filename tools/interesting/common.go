package interesting

import (
	"fmt"
	"strings"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

/*

- O que significa esses parâmetros ?!
	• "HTTP Index Of's" =  Esse endereço guarda as URL's das páginas que contém Index Of.

*/

// DirectoryPlugins :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: Database.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: Database.OtherInformationsSlice["target.http.indexof"]
func DirectoryPlugins() models.InformationsModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/plugins/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/plugins Index Of Raw", response.Raw)
	}

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

func DirectoryUploads() models.InformationsModel {
	var http = net.NewNETClient().SetURL(database.Memory.GetString("Target"))
	http.SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/uploads/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/uploads Index Of Raw", response.Raw)
	}

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

func DirectoryThemes() models.InformationsModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/themes/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/themes Index Of Raw", response.Raw)
	}

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

// AdminPage :: Simple requests to see if there is.
func AdminPage() models.InformationsModel {
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
	
	var model = models.InformationsModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}
	
	defer database.Memory.SetString("HTTP Admin Page Status", response.Response.Status)
	
	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

func Robots() models.InformationsModel {
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

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.sitemap.xml.status"]
func Sitemap() models.InformationsModel {
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

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}

func Readme() models.InformationsModel {
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

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		model.Confidence = 100
	}

	return model
}


func XMLRPC() models.InformationsModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("xmlrpc.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	var model = models.InformationsModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if strings.Contains(response.Raw, "XML-RPC server accepts POST requests only.") {
		model.Confidence += 60
	}

	if response.Response.StatusCode == 405 {
		model.Confidence += 20
	}
	
	if strings.Contains(response.Raw, "This error was generated by Mod_Security.") {
		model.Confidence = 80
		model.FoundBy = "Direct Access (Mod_Security)"
	}
	
	if strings.Contains(database.Memory.GetString("HTTP Index Raw"), "xmlrpc.php") {
		model.Confidence += 11
	}

	return model
}