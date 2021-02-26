package commons

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
)

func DirectoryUploads() *gohttp.Response {
	http := gohttp.NewHTTPClient().SetURL(database.Memory.GetString("Target"))
	http.SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/uploads/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/uploads Index Of Raw", response.Raw)
	}

	return response
}

// DirectoryPlugins :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: Database.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: Database.OtherInformationsSlice["target.http.indexof"]
func DirectoryPlugins() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/plugins/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/plugins Index Of Raw", response.Raw)
	}

	return response
}

// DirectoryThemes :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: Database.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: Database.OtherInformationsSlice["target.http.indexof"]
func DirectoryThemes() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/themes/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/themes Index Of Raw", response.Raw)
	}

	return response
}

// AdminPage :: Simple requests to see if there is.
func AdminPage() (string, *gohttp.Response) {
	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("wp-admin/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	switch response.Response.StatusCode {
	case 200:
		database.Memory.SetString("HTTP Admin Page Status", "true")

		return "true", response
	case 403:
		database.Memory.SetString("HTTP Admin Page Status", "redirect")

		return "redirect", response
	default:
		return "false", response
	}
}

// Robots :: Simple requests to see if there is.
// The command's message will be saved on this map :: Database.OtherInformationsString["target.http.robots.txt.status"]
// The source code of the robots file will be saved within this map :: Database.OtherInformationsString["target.http.robots.txt.raw"]
func Robots() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("robots.txt")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	return response
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.sitemap.xml.status"]
func Sitemap() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("sitemap.xml")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	return response
}

func Readme() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("readme.html")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	return response
}

// XMLRPC :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.xmlrpc.php.status"]
func XMLRPC() (string, string) {
	if strings.Contains(database.Memory.GetString("HTTP Index Raw"), "xmlrpc.php") {
		return "", "Link tag"
	}

	http := gohttp.NewHTTPClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("xmlrpc.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	// Status Code Return: 405
	if strings.Contains(response.Raw, "XML-RPC server accepts POST requests only.") {
		return "Confirmed", "Direct Access"
	} else if strings.Contains(response.Raw, "This error was generated by Mod_Security.") {
		return "Mod_Security", "Direct Access"
	} else if response.Response.StatusCode == 403 {
		return "Forbidden", "Direct Access"
	}

	return "False", ""

}
