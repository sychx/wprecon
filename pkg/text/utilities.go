package text

import (
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html/charset"
)

// Decode :: This function will convert any string to UTF-8.
func Decode(text string) string {
	r, err := charset.NewReader(strings.NewReader(text), "latin1")
	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}
