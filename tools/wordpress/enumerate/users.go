package enumerate

import (
	"encoding/json"
	"fmt"
	"regexp"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/text"
)

type uJSON []struct {
	Name string `json:"name"`
}

// UsersEnumeratePassive :: Enumerate using feed
func UsersEnumeratePassive() ([]string, *gohttp.Response) {
	response := gohttp.SimpleRequest(Database.Target, "feed/")

	rex := regexp.MustCompile("<dc:creator><!\\[CDATA\\[(.+?)\\]\\]></dc:creator>")
	submatch := rex.FindAllSubmatch([]byte(response.Raw), -1)

	for _, value := range submatch {
		valueString := fmt.Sprintf("%s", value[1])

		if _, has := text.ContainsSliceString(Database.OtherInformationsSlice["target.http.users"], valueString); !has {
			Database.OtherInformationsSlice["target.http.users"] = append(Database.OtherInformationsSlice["target.http.users"], valueString)
		}
	}

	if len(Database.OtherInformationsSlice["target.http.users"]) > 0 {
		Database.OtherInformationsString["target.http.users.method"] = "Feed"
	}

	return Database.OtherInformationsSlice["target.http.users"], response
}

// UsersEnumerateAgressive ::
func UsersEnumerateAgressive() ([]string, *gohttp.Response) {
	var responseReturn *gohttp.Response
	var ujson uJSON
	done := false

	// Enumerate using Yoast SEO
	func() {
		if done == false {
			response := gohttp.SimpleRequest(Database.Target, "author-sitemap.xml")

			rex := regexp.MustCompile("<loc>.*?/author/(.*?)/</loc>")

			submatch := rex.FindAllSubmatch([]byte(response.Raw), -1)

			for _, value := range submatch {
				valueString := fmt.Sprintf("%s", value[1])

				if _, has := text.ContainsSliceString(Database.OtherInformationsSlice["target.http.users"], valueString); !has {
					Database.OtherInformationsSlice["target.http.users"] = append(Database.OtherInformationsSlice["target.http.users"], valueString)
				}
			}

			if len(Database.OtherInformationsSlice["target.http.users"]) > 0 {
				Database.OtherInformationsString["target.http.users.method"] = "YoastSEO"
				done = true
			}

			responseReturn = response
		}
	}()

	// Enumerate using route
	func() {
		if done == false {
			response := gohttp.SimpleRequest(Database.Target, "?rest_route=/wp/v2/users")

			if response.Response.StatusCode == 200 && response.Raw != "" {
				json.Unmarshal([]byte(response.Raw), &ujson)

				for _, value := range ujson {
					if _, has := text.ContainsSliceString(Database.OtherInformationsSlice["target.http.users"], value.Name); !has {
						Database.OtherInformationsSlice["target.http.users"] = append(Database.OtherInformationsSlice["target.http.users"], value.Name)
					}
				}

				if len(Database.OtherInformationsSlice["target.http.users"]) > 0 {
					Database.OtherInformationsString["target.http.users.method"] = "Route"
					done = true
				}
			} else if response.Response.StatusCode == 401 && response.Raw != "" && Database.Verbose {
				printer.Danger("Status code 401, I don't think I'm allowed to list users. Target Url:", response.URL.Full, "— Target source code:", response.Raw)
			}

			responseReturn = response
		}
	}()

	// Enumerate using json file
	func() {
		if done == false {
			response := gohttp.SimpleRequest(Database.Target, "wp-json/wp/v2/users")

			if response.Response.StatusCode == 200 && response.Raw != "" {
				json.Unmarshal([]byte(response.Raw), &ujson)

				for _, value := range ujson {
					if _, has := text.ContainsSliceString(Database.OtherInformationsSlice["target.http.users"], value.Name); !has {
						Database.OtherInformationsSlice["target.http.users"] = append(Database.OtherInformationsSlice["target.http.users"], value.Name)
					}
				}

				if len(Database.OtherInformationsSlice["target.http.users"]) > 0 {
					Database.OtherInformationsString["target.http.users.method"] = "Json"
					done = true
				}
			} else if response.Response.StatusCode == 401 && response.Raw != "" && Database.Verbose {
				printer.Danger("Status code 401, I don't think I'm allowed to list users. Target Url:", response.URL.Full, "— Target source code:", response.Raw)
			}

			responseReturn = response
		}
	}()

	return Database.OtherInformationsSlice["target.http.users"], responseReturn
}
