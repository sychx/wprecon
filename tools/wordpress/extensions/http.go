package extensions

import (
	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

var http = gohttp.HTTPOptions{URL: gohttp.URLOptions{}, Options: gohttp.Options{}}

// SimpleRequest :: As the name says a simple request, with almost all the parameters of a common request with gohttp.
func SimpleRequest(url string, directory string) *gohttp.Response {
	http.URL.Simple = url
	http.URL.Directory = directory
	http.Options.TLSCertificateVerify = InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"]
	http.Options.Tor = InfosWprecon.OtherInformationsBool["http.options.tor"]
	http.Options.RandomUserAgent = InfosWprecon.OtherInformationsBool["http.options.randomuseragent"]

	response, err := gohttp.HTTPRequest(&http)

	if err != nil {
		printer.Fatal(err)
	}

	return &response
}
