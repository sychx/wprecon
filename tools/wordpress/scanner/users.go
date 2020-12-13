package wpscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

/*
Here you may feel somewhat confused. But don't be surprised that I explain it to you!
As each function returns a different value, I chose to do it that way. But in the future I will improve this.
*/

type usersJson []struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type usersRe struct {
	Name string
	Slug string
}

type Users struct {
	Verbose bool
	Request gohttp.Http
}

func (options *Users) EnumerateJson() (bool, usersJson) {
	/* Start of the first scan */
	options.Request.Dir = "wp-json/wp/v2/users"

	switch response, err := gohttp.HttpRequest(options.Request); true {
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
	options.Request.Dir = "?rest_route=/wp/v2/users"

	switch response, err := gohttp.HttpRequest(options.Request); true {
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

	return false, nil
}

func (options *Users) EnumerateRss() (bool, []usersRe) {
	options.Request.Dir = "feed/"

	switch response, err := gohttp.HttpRequest(options.Request); true {
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
