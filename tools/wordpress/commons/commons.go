package commons

import (
	"strings"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

var http = gohttp.HTTPOptions{URL: gohttp.URLOptions{}, Options: gohttp.Options{}}

// DirectoryPlugins :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: InfosWprecon.OtherInformationsSlice["target.http.indexof"]
func DirectoryPlugins() *gohttp.Response {
	http.URL.Simple = InfosWprecon.Target
	http.URL.Directory = "/wp-content/plugins/"
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = true

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	if strings.Contains(response.Raw, "Index of") {
		InfosWprecon.OtherInformationsSlice["target.http.indexof"] = append(InfosWprecon.OtherInformationsSlice["target.http.indexof"], http.URL.Full)
		InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"] = response.Raw

		if InfosWprecon.Verbose {
			printer.Warning("\"index of\" found, in", http.URL.Full)
		}
	}

	return &response
}

// DirectoryThemes :: Simple requests to see if it exists and if it has index of.
// If this directory is identified with Index Of, its source code will be saved in this map :: InfosWprecon.OtherInformationsString["target.http.wp-content/plugin.indexof.raw"]
// Any directory that is identified with Index Of will be saved on this map :: InfosWprecon.OtherInformationsSlice["target.http.indexof"]
func DirectoryThemes() *gohttp.Response {
	http.URL.Simple = InfosWprecon.Target
	http.URL.Directory = "/wp-content/themes/"
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = true

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	if strings.Contains(response.Raw, "Index of") {
		InfosWprecon.OtherInformationsSlice["target.http.indexof"] = append(InfosWprecon.OtherInformationsSlice["target.http.indexof"], http.URL.Full)
		InfosWprecon.OtherInformationsString["target.http.wp-content/themes.indexof.raw"] = response.Raw

		if InfosWprecon.Verbose {
			printer.Warning("\"index of\" found, in", http.URL.Full)
		}
	}

	return &response
}

// AdminPage :: Simple requests to see if there is.
func AdminPage() (string, *gohttp.Response) {
	http.URL.Simple = InfosWprecon.Target
	http.URL.Directory = "wp-admin/"
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = InfosWprecon.OtherInformationsBool["http.options.randomuseragent"]

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 {
		return "true", &response
	} else if response.StatusCode == 302 {
		return "redirect", &response
	} else {
		return "false", &response
	}
}

// Robots :: Simple requests to see if there is.
// The command's message will be saved on this map :: InfosWprecon.OtherInformationsString["target.http.robots.txt.status"]
// The source code of the robots file will be saved within this map :: InfosWprecon.OtherInformationsString["target.http.robots.txt.raw"]
func Robots() *gohttp.Response {
	http.URL.Simple = InfosWprecon.Target
	http.URL.Directory = "robots.txt"
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = InfosWprecon.OtherInformationsBool["http.options.randomuseragent"]

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 {
		InfosWprecon.OtherInformationsString["target.http.robots.txt.raw"] = response.Raw
		InfosWprecon.OtherInformationsString["target.http.robots.txt.status"] = "sucess"

		if InfosWprecon.Verbose && response.Raw != "" {
			printer.Warning("Robots.txt file text:")
			printer.Println(response.Raw)
		}
	}

	return &response
}

// Sitemap :: Simple requests to see if there is.
// The command's message will be saved on this map. :: InfosWprecon.OtherInformationsString["target.http.sitemap.xml.status"]
func Sitemap() *gohttp.Response {
	http.URL.Simple = InfosWprecon.Target
	http.URL.Directory = "sitemap.xml"
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = InfosWprecon.OtherInformationsBool["http.options.randomuseragent"]

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	if response.StatusCode == 200 {
		InfosWprecon.OtherInformationsString["target.http.sitemap.xml.status"] = "true"

		if InfosWprecon.Verbose {
			printer.Warning("Sitemap.xml found:", response.URL.Full)
		}
	}

	return &response
}

// XMLRPC :: Simple requests to see if there is.
// The command's message will be saved on this map. :: InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"]
func XMLRPC() (string, *gohttp.Response) {
	http.URL.Simple = InfosWprecon.Target
	http.URL.Directory = "xmlrpc.php"
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = InfosWprecon.OtherInformationsBool["http.options.randomuseragent"]

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	// Status Code Return: 405
	if strings.Contains(response.Raw, "XML-RPC server accepts POST requests only.") {
		InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"] = "Sucess"

		return "true", &response
	} else if strings.Contains(response.Raw, "This error was generated by Mod_Security.") {
		InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"] = "Mod_Security"

		return "mod_security", &response
	} else if response.StatusCode == 403 {
		InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.status"] = "Forbidden"

		return "forbidden", &response
	}

	return "false", &response
}
