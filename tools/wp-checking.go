package tools

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/blkzy/wpsgo/pkg/gohttp"
)

func Wordpress(response gohttp.Response) (string, error) {
	data, _ := ioutil.ReadAll(response.Body)

	payloads := [...]string{
		`<meta name="generator content="WordPress`,
		`<a href="http://www.wordpress.com">Powered by WordPress</a>`,
		`<link rel='https://api.w.org/'`}

	exists := 0

	for _, value := range payloads {
		if strings.Contains(value, string(data)) {
			exists++
		}
	}

	var calc = exists / 3 * 100

	return fmt.Sprintf("%v", calc), nil
}
