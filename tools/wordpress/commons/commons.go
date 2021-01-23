package commons

import (
	"fmt"
	"strings"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// DirectoryPlugins :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: InfosWprecon.OtherInformationsSlice["target.http.indexof"]
func DirectoryPlugins() *gohttp.Response {
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target).SetURLDirectory("/wp-content/plugins/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprint(err))
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
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target).SetURLDirectory("/wp-content/themes/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprint(err))
	}

	if strings.Contains(response.Raw, "Index of") {
		InfosWprecon.OtherInformationsSlice["target.http.indexof"] = append(InfosWprecon.OtherInformationsSlice["target.http.indexof"], response.URL.Full)
		InfosWprecon.OtherInformationsString["target.http.wp-content/themes.indexof.raw"] = response.Raw
	}

	return response
}

// AdminPage :: Simple requests to see if there is.
func AdminPage() (string, *gohttp.Response) {
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target).SetURLDirectory("wp-admin/")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprint(err))
	}

	if response.Response.StatusCode == 200 {
		return "true", response
	} else if response.Response.StatusCode == 302 {
		return "redirect", response
	} else {
		return "false", response
	}
}

// Robots :: Simple requests to see if there is.
// The command's message will be saved on this map :: InfosWprecon.OtherInformationsString["target.http.robots.txt.status"]
// The source code of the robots file will be saved within this map :: InfosWprecon.OtherInformationsString["target.http.robots.txt.raw"]
func Robots() *gohttp.Response {
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target).SetURLDirectory("robots.txt")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprint(err))
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
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target).SetURLDirectory("sitemap.xml")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprint(err))
	}

	if response.Response.StatusCode == 200 {
		InfosWprecon.OtherInformationsString["target.http.sitemap.xml.status"] = "true"
	}

	return response
}

// XMLRPC :: Simple requests to see if there is.
// The command's message will be saved on this map. :: InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"]
func XMLRPC() (string, *gohttp.Response) {
	http := gohttp.NewHTTPClient().SetURL(InfosWprecon.Target).SetURLDirectory("xmlrpc.php")
	http.OnTor(InfosWprecon.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(InfosWprecon.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"])

	response, err := http.Run()

	if err != nil {
		printer.Danger(fmt.Sprint(err))
	}

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
