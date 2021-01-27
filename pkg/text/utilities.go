package text

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/blackbinn/wprecon/pkg/printer"
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

func GetOneImportantFile(raw string) string {
	rex := regexp.MustCompile("<a href=\"(readme*?|README*?|Readme*?|Changelog*?|changelog*?|CHANGELOG*?)\">.*?</a>")

	submatchall := rex.FindStringSubmatch(raw)

	if len(submatchall) > 0 {
		return submatchall[1]
	}

	return fmt.Sprint(submatchall)
}

// GetFileExtensions :: This function searches for files by their extension, within an index of.
func GetFileExtensions(url string, raw string) [][][]byte {
	//rex := regexp.MustCompile("<a href=\"(.*?.sql|.*?.zip|.*?.tar|.*?.tar.gz)\">.*?</a>")

	rex := regexp.MustCompile("<a href=\"(.*?.sql|.*?.zip|.*?.tar|.*?.tar.gz)\">.*?</a>")
	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	return submatchall
}

// GetVersionStableTag :: This function searches for the version of the plugin or theme.
func GetVersionStableTag(raw string) string {
	rex := regexp.MustCompile("(?:Stable tag:|stable tag:|Version:|version:|version) ([0-9.-]+)")

	submatchall := rex.FindSubmatch([]byte(raw))

	if len(submatchall) > 0 {
		version := fmt.Sprintf("%s", submatchall[1])

		return version
	}

	return ""
}

// GetVersionChangelog :: This function searches for the version of the plugin or theme.
func GetVersionChangelog(raw string) string {
	rex := regexp.MustCompile("=+\\s+(?:v(?:ersion)?\\s*)?([0-9.-]+)[ \ta-z0-9().\\-/]*=+")

	submatchall := rex.FindSubmatch([]byte(raw))

	if len(submatchall) > 0 {
		version := fmt.Sprintf("%s", submatchall[1])

		return version
	}

	return ""
}

func ReadAllFile(filename string) (chars []string, count int) {
	file, err := os.Open(filename)

	if err != nil {
		printer.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chars = append(chars, scanner.Text())
	}

	return chars, len(chars)
}
