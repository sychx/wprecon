package commons

import (
	"fmt"
	"strings"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
)

func DirectoryUploads() *gohttp.Response {
	http := gohttp.NewHTTPClient().SetURL(Database.Target)
	http.SetURLDirectory(Database.WPContent + "/uploads/")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		Database.OtherInformationsSlice["target.http.indexof"] = append(Database.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		Database.OtherInformationsString["target.http.wp-content/uploads.indexof.raw"] = response.Raw
	}

	return response
}

// DirectoryPlugins :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: Database.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: Database.OtherInformationsSlice["target.http.indexof"]
func DirectoryPlugins() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(Database.Target).SetURLDirectory(Database.WPContent + "/plugins/")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		Database.OtherInformationsSlice["target.http.indexof"] = append(Database.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		Database.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"] = response.Raw
	}

	return response
}

// DirectoryThemes :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: Database.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: Database.OtherInformationsSlice["target.http.indexof"]
func DirectoryThemes() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(Database.Target).SetURLDirectory(Database.WPContent + "/themes/")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		Database.OtherInformationsSlice["target.http.indexof"] = append(Database.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		Database.OtherInformationsString["target.http.wp-content/themes.indexof.raw"] = response.Raw
	}

	return response
}

// AdminPage :: Simple requests to see if there is.
func AdminPage() (string, *gohttp.Response) {
	http := gohttp.NewHTTPClient()
	http.SetURL(Database.Target).SetURLDirectory("wp-admin/")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	switch response.Response.StatusCode {
	case 200:
		Database.OtherInformationsString["target.http.admin-page"] = response.URL.Full
		return "true", response
	case 403:
		Database.OtherInformationsString["target.http.admin-page"] = response.URL.Full
		return "redirect", response
	default:
		Database.OtherInformationsString["target.http.admin-page"] = ""
		return "false", response
	}
}

// Robots :: Simple requests to see if there is.
// The command's message will be saved on this map :: Database.OtherInformationsString["target.http.robots.txt.status"]
// The source code of the robots file will be saved within this map :: Database.OtherInformationsString["target.http.robots.txt.raw"]
func Robots() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(Database.Target).SetURLDirectory("robots.txt")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if response.Response.StatusCode == 200 {
		Database.OtherInformationsString["target.http.robots.txt.raw"] = response.Raw
		Database.OtherInformationsString["target.http.robots.txt.status"] = "sucess"
	}

	return response
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.sitemap.xml.status"]
func Sitemap() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(Database.Target).SetURLDirectory("sitemap.xml")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if response.Response.StatusCode == 200 {
		Database.OtherInformationsString["target.http.sitemap.xml.status"] = "true"
	}

	return response
}

func Readme() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(Database.Target).SetURLDirectory("readme.html")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	return response
}

// XMLRPC :: Simple requests to see if there is.
// The command's message will be saved on this map. :: Database.OtherInformationsString["target.http.xmlrpc.php.status"]
func XMLRPC() (string, *gohttp.Response) {
	if strings.Contains(Database.OtherInformationsString["target.http.index.raw"], "xmlrpc.php") {
		Database.OtherInformationsString["target.http.xmlrpc.php.checkedby"] = "Link tag"

		return "Link tag", &gohttp.Response{}
	} else {
		http := gohttp.NewHTTPClient()
		http.SetURL(Database.Target).SetURLDirectory("xmlrpc.php")
		http.OnTor(Database.OtherInformationsBool["http.options.tor"])
		http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
		http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
		http.FirewallDetection(true)

		response, err := http.Run()

		if err != nil {
			printer.Danger(fmt.Sprintf("%s", err))
		}

		Database.OtherInformationsString["target.http.xmlrpc.php.checkedby"] = "Access"

		// Status Code Return: 405
		if strings.Contains(response.Raw, "XML-RPC server accepts POST requests only.") {
			Database.OtherInformationsString["target.http.xmlrpc.php.status"] = "Sucess"

			return "true", response
		} else if strings.Contains(response.Raw, "This error was generated by Mod_Security.") {
			Database.OtherInformationsString["target.http.xmlrpc.php.status"] = "Mod_Security"

			return "mod_security", response
		} else if response.Response.StatusCode == 403 {
			Database.OtherInformationsString["target.http.xmlrpc.php.status"] = "Forbidden"

			return "forbidden", response
		}
		return "false", response
	}
}
