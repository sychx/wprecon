package enumerate

import (
	"encoding/json"
	"regexp"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/text"
)

type uJSON []struct {
	Name string `json:"name"`
}

// UsersEnumeratePassive :: Enumerate using feed
func UsersEnumeratePassive() (users []string, method string, URL string) {
	response := gohttp.SimpleRequest(database.Memory.GetString("Target"), "feed/")

	rex := regexp.MustCompile("<dc:creator><!\\[CDATA\\[(.+?)\\]\\]></dc:creator>")

	for _, user := range rex.FindAllStringSubmatch(response.Raw, -1) {
		if _, has := text.ContainsSliceString(users, user[1]); !has {
			users = append(users, user[1])
		}
	}

	URL = response.URL.Full
	method = "Feed"

	return
}

// UsersEnumerateAgressive :: In its aggressive mode, wprecon tries to enumerate users using 4 types of enumeration, which can be considered difficult to access for an ordinary user, and for this reason they are classified as aggressive enumeration.
func UsersEnumerateAgressive() (users []string, method string, URL string) {
	var ujson uJSON
	var done bool

	// Enumerate using route
	func() {
		if done == false {
			response := gohttp.SimpleRequest(database.Memory.GetString("Target"), "?rest_route=/wp/v2/users")

			if response.Response.StatusCode == 200 && response.Raw != "" {
				json.Unmarshal([]byte(response.Raw), &ujson)

				for _, value := range ujson {
					if _, has := text.ContainsSliceString(users, value.Name); !has {
						users = append(users, value.Name)
						done = true
					}
				}

				method = "Route API"
				URL = response.URL.Full

				return
			} else if response.Response.StatusCode == 401 && response.Raw != "" {
				printer.Danger("Status code 401, I don't think I'm allowed to list users. Target Url:", response.URL.Full, "— Target source code:", response.Raw).L()
			}
		}
	}()

	// Enumerate using Yoast SEO
	func() {
		if done == false {
			response := gohttp.SimpleRequest(database.Memory.GetString("Target"), "author-sitemap.xml")

			if response.Response.StatusCode == 200 && response.Raw != "" {
				rex := regexp.MustCompile("<loc>.*?/author/(.*?)/</loc>")

				for _, value := range rex.FindAllStringSubmatch(response.Raw, -1) {
					if _, has := text.ContainsSliceString(users, value[1]); !has {
						users = append(users, value[1])
						done = true
					}
				}

				URL = response.URL.Full
				method = "Yoast SEO"

				return
			}
		}
	}()

	// Enumerate using json file
	func() {
		if done == false {
			response := gohttp.SimpleRequest(database.Memory.GetString("Target"), "wp-json/wp/v2/users")

			if response.Response.StatusCode == 200 && response.Raw != "" {
				json.Unmarshal([]byte(response.Raw), &ujson)

				for _, value := range ujson {
					if _, has := text.ContainsSliceString(users, value.Name); !has {
						users = append(users, value.Name)
						done = true
					}
				}

				URL = response.URL.Full
				method = "Route Json API"

				return
			} else if response.Response.StatusCode == 401 && response.Raw != "" {
				printer.Danger("Status code 401, I don't think I'm allowed to list users. Target Url:", response.URL.Full, "— Target source code:", response.Raw).L()
			}

		}
	}()

	users, method, URL = UsersEnumeratePassive()

	return
}
