package commons

import (
	"fmt"
	"strings"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
)

func DirectoryUploads() *gohttp.Response {
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target)
	http.SetURLDirectory(InfosWprecon.WPContent + "/uploads/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		InfosWprecon.OtherInformationsSlice["target.http.indexof"] = append(InfosWprecon.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		InfosWprecon.OtherInformationsString["target.http.wp-content/uploads.indexof.raw"] = response.Raw
	}

	return response
}

// DirectoryPlugins :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: InfosWprecon.OtherInformationsSlice["target.http.indexof"]
func DirectoryPlugins() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(InfosWprecon.Target).SetURLDirectory(InfosWprecon.WPContent + "/plugins/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		InfosWprecon.OtherInformationsSlice["target.http.indexof"] = append(InfosWprecon.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"] = response.Raw
	}

	return response
}

// DirectoryThemes :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: InfosWprecon.OtherInformationsSlice["target.http.indexof"]
func DirectoryThemes() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(InfosWprecon.Target).SetURLDirectory(InfosWprecon.WPContent + "/themes/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if strings.Contains(response.Raw, "Index of") {
		InfosWprecon.OtherInformationsSlice["target.http.indexof"] = append(InfosWprecon.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		InfosWprecon.OtherInformationsString["target.http.wp-content/themes.indexof.raw"] = response.Raw
	}

	return response
}

// AdminPage :: Simple requests to see if there is.
func AdminPage() (string, *gohttp.Response) {
	http := gohttp.NewHTTPClient()
	http.SetURL(InfosWprecon.Target).SetURLDirectory("wp-admin/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	switch response.Response.StatusCode {
	case 200:
		InfosWprecon.OtherInformationsString["target.http.admin-page"] = response.URL.Full
		return "true", response
	case 403:
		InfosWprecon.OtherInformationsString["target.http.admin-page"] = response.URL.Full
		return "redirect", response
	default:
		InfosWprecon.OtherInformationsString["target.http.admin-page"] = ""
		return "false", response
	}
}

// Robots :: Simple requests to see if there is.
// The command's message will be saved on this map :: InfosWprecon.OtherInformationsString["target.http.robots.txt.status"]
// The source code of the robots file will be saved within this map :: InfosWprecon.OtherInformationsString["target.http.robots.txt.raw"]
func Robots() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(InfosWprecon.Target).SetURLDirectory("robots.txt")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if response.Response.StatusCode == 200 {
		InfosWprecon.OtherInformationsString["target.http.robots.txt.raw"] = response.Raw
		InfosWprecon.OtherInformationsString["target.http.robots.txt.status"] = "sucess"
	}

	return response
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: InfosWprecon.OtherInformationsString["target.http.sitemap.xml.status"]
func Sitemap() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(InfosWprecon.Target).SetURLDirectory("sitemap.xml")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if response.Response.StatusCode == 200 {
		InfosWprecon.OtherInformationsString["target.http.sitemap.xml.status"] = "true"
	}

	return response
}

func Readme() *gohttp.Response {
	http := gohttp.NewHTTPClient()
	http.SetURL(InfosWprecon.Target).SetURLDirectory("readme.html")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	return response
}

// XMLRPC :: Simple requests to see if there is.
// The command's message will be saved on this map. :: InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"]
func XMLRPC() (string, *gohttp.Response) {
	if strings.Contains(InfosWprecon.OtherInformationsString["target.http.index.raw"], "xmlrpc.php") {
		InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.checkedby"] = "Link tag"

		return "Link tag", &gohttp.Response{}
	} else {
		http := gohttp.NewHTTPClient()
		http.SetURL(InfosWprecon.Target).SetURLDirectory("xmlrpc.php")
		http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
		http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
		http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])
		http.FirewallDetection(true)

		response, err := http.Run()

		if err != nil {
			printer.Danger(fmt.Sprintf("%s", err))
		}

		InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.checkedby"] = "Access"

		// Status Code Return: 405
		if strings.Contains(response.Raw, "XML-RPC server accepts POST requests only.") {
			InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"] = "Sucess"

			return "true", response
		} else if strings.Contains(response.Raw, "This error was generated by Mod_Security.") {
			InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"] = "Mod_Security"

			return "mod_security", response
		} else if response.Response.StatusCode == 403 {
			InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"] = "Forbidden"

			return "forbidden", response
		}
		return "false", response
	}
}
