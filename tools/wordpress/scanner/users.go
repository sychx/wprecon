package wpscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/text"
)

/*
Here you may feel somewhat confused. But don't be surprised that I explain it to you!
As each function returns a different value, I chose to do it that way. But in the future I will improve this.
*/

type Users struct {
	Verbose bool
	Request gohttp.Http
}

type uJson []struct {
	Name string `json:"name"`
}

func (options *Users) Enumerate() (bool, []string) {

	if has, json := options.json(); has {
		names := make([]string, len(json))

		for key, value := range json {

			names[key] = fmt.Sprintf("%s", value.Name)
		}

		return true, names

	} else if has, route := options.route(); has {
		names := make([]string, len(route))

		for key, value := range route {

			names[key] = fmt.Sprintf("%s", value.Name)
		}

		return true, names

	} else if has, rss := options.rss(); has {
		var names []string

		for _, value := range rss {

			valueString := fmt.Sprintf("%s", value[1])

			if _, has := text.ContainsSliceString(names, valueString); !has && valueString != "" {
				names = append(names, valueString)
			}
		}

		return true, names
	}

	return false, nil
}

func (options *Users) json() (bool, uJson) {
	options.Request.Dir = "wp-json/wp/v2/users"

	var jsn uJson

	if response, err := gohttp.HttpRequest(options.Request); response.StatusCode == 200 {

		json.NewDecoder(response.Body).Decode(&jsn)

		if len(jsn) > 0 {
			return true, jsn
		}

	} else if err != nil {
		printer.Fatal(err)
	}

	return false, nil
}

func (options *Users) route() (bool, uJson) {
	options.Request.Dir = "?rest_route=/wp/v2/users"

	var jsn uJson

	if response, err := gohttp.HttpRequest(options.Request); response.StatusCode == 200 {

		json.NewDecoder(response.Body).Decode(&jsn)

		if len(jsn) > 0 {
			return true, jsn
		}

	} else if err != nil {
		printer.Fatal(err)
	}

	return false, nil
}

func (options *Users) rss() (bool, [][][]byte) {
	options.Request.Dir = "feed/"

	if response, err := gohttp.HttpRequest(options.Request); response.StatusCode == 200 {
		re := regexp.MustCompile("<dc:creator><!\\[CDATA\\[(.+?)\\]\\]></dc:creator>")

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printer.Fatal(err)
		}

		submatchall := re.FindAllSubmatch([]byte(bodyBytes), -1)

		return true, submatchall

	} else if err != nil {
		printer.Fatal(err)
	}

	return false, nil
}
