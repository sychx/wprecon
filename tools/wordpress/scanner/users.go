package scanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/text"
)

// Users ::
type Users struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
}

type uJSON []struct {
	Name string `json:"name"`
}

// Enumerate ::
func (options *Users) Enumerate() {
	if has, json := options.json(); has {
		printer.Done("⎡ User(s) :")

		for _, value := range json {
			printer.Done("⎢", value.Name)
		}
	} else if has, route := options.route(); has {
		printer.Done("⎡ User(s) :")

		for _, value := range route {
			printer.Done("⎢", value.Name)
		}
	} else if has, rss := options.rss(); has {
		var names []string

		printer.Done("⎡ User(s) :")

		for _, value := range rss {
			valueString := fmt.Sprintf("%s", value[1])

			if _, has := text.ContainsSliceString(names, valueString); !has && valueString != "" {
				printer.Done("⎢", valueString)
			}
		}
	} else {
		printer.Danger("Unfortunately no user was found. ;-;")
	}

	printer.Println("")
}

func (options *Users) json() (bool, uJSON) {
	options.HTTP.URL.Directory = "wp-json/wp/v2/users"

	var jsn uJSON

	if response, err := gohttp.HTTPRequest(options.HTTP); response.StatusCode == 200 {

		json.NewDecoder(response.Raw).Decode(&jsn)

		if len(jsn) > 0 {
			return true, jsn
		}

	} else if err != nil {
		printer.Fatal(err)
	}

	return false, nil
}

func (options *Users) route() (bool, uJSON) {
	options.HTTP.URL.Directory = "?rest_route=/wp/v2/users"

	var jsn uJSON

	if response, err := gohttp.HTTPRequest(options.HTTP); response.StatusCode == 200 {

		json.NewDecoder(response.Raw).Decode(&jsn)

		if len(jsn) > 0 {
			return true, jsn
		}

	} else if err != nil {
		printer.Fatal(err)
	}

	return false, nil
}

func (options *Users) rss() (bool, [][][]byte) {
	options.HTTP.URL.Directory = "feed/"

	if response, err := gohttp.HTTPRequest(options.HTTP); response.StatusCode == 200 {
		re := regexp.MustCompile("<dc:creator><!\\[CDATA\\[(.+?)\\]\\]></dc:creator>")

		bodyBytes, err := ioutil.ReadAll(response.Raw)

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
