package wpscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

type usersJson []struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type usersRe struct {
	Name string
	Slug string
}

func UserEnum(target string, randomUserAgent bool) {
	printer.Loading("Hunting users...")

	if hasEnum, users := userEnumJson(target, randomUserAgent); hasEnum == true {
		for _, user := range users {
			printer.LoadingDone("User:", user.Name, "—", "Slug:", user.Slug)
		}
	} else if hasEnum, users := userEnumRss(target, randomUserAgent); hasEnum == true {
		for _, user := range users {
			if user.Name != "" {
				printer.LoadingDone("User:", user.Name)
			}
		}
	} else {
		printer.LoadingDanger("Unfortunately no user was found. ;-;")
	}
}

func userEnumJson(target string, randomUserAgent bool) (bool, usersJson) {

	/* Start of the first scan */
	switch response, err := gohttp.HttpRequest(gohttp.Http{URL: target, Dir: "wp-json/wp/v2/users", RandomUserAgent: randomUserAgent}); true {
	case response.StatusCode == 200:
		var jsn usersJson
		json.NewDecoder(response.Body).Decode(&jsn)

		if len(jsn) > 0 {
			return true, jsn
		}

	case err != nil:
		printer.Fatal(err)
	}
	/* End of first scan */

	/* Start of the second check */
	switch response, err := gohttp.HttpRequest(gohttp.Http{URL: target, Dir: "?rest_route=/wp/v2/users", RandomUserAgent: randomUserAgent}); true {
	case response.StatusCode == 200:
		var jsn usersJson
		json.NewDecoder(response.Body).Decode(&jsn)

		if len(jsn) > 0 {
			return true, jsn
		}

	case err != nil:
		printer.Fatal(err)
	}
	/* End of second check */

	/* Start of the third check */
	/* End of third check */

	return false, nil
}

func userEnumRss(target string, randomUserAgent bool) (bool, []usersRe) {
	switch response, err := gohttp.HttpRequest(gohttp.Http{URL: target, Dir: "feed/", RandomUserAgent: randomUserAgent}); true {
	case response.StatusCode == 200:
		re := regexp.MustCompile("<dc:creator><!\\[CDATA\\[(.+?)\\]\\]></dc:creator>")

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		submatchall := re.FindAllSubmatch([]byte(bodyBytes), -1)

		for key, name := range submatchall {
			dir := make([]usersRe, len(submatchall))

			dir[key].Name = fmt.Sprintf("%s", name[1])

			return true, dir
		}

	case err != nil:
		printer.Fatal(err)
	}

	return false, nil
}
