package text

import (
	"bufio"
	"encoding/csv"
	"os"
	"regexp"

	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

const (
	MatchImportantFile     = "<a href=\"([[[R|r]eadme|EADME]|[C|c]hangelog|HANGELOG]|[R|r]elease_log].*?)\">.*?</a>"
	MatchFileExtensions    = "<a href=\"(.*?.[sql|db|zip|tar|tar.gz])\">.*?</a>"
	MatchVersionStableTag  = "[S|s]table [T|t]ag.*?([0-9.-]+)"
	MatchVersionChangelog  = "=+\\s+(?:v(?:ersion)?\\s*)?([0-9.-]+)[ \ta-z0-9().\\-/]*=+"
	MatchVersionReleaseLog = "[v|V]ersion.*?([0-9.-]+)"
)

// GetOneImportantFile :: This function will do a search for the source code to search for an important file.
func GetOneImportantFile(raw string) string {
	re := regexp.MustCompile(MatchImportantFile)

	submatchall := re.FindStringSubmatch(raw)

	if len(submatchall) > 0 {
		return submatchall[1]
	}

	return ""
}

// GetFileExtensions :: This function searches for files by their extension, within an index of.
func GetFileExtensions(url string, raw string) [][][]byte {
	re := regexp.MustCompile(MatchFileExtensions)
	submatchall := re.FindAllSubmatch([]byte(raw), -1)

	return submatchall
}

// GetVersionStableTag :: This function searches for the version of the plugin or theme.
func GetVersionStableTag(raw string) []string {
	re := regexp.MustCompile(MatchVersionStableTag)

	return re.FindStringSubmatch(raw)
}

// GetVersionChangelog :: This function searches for the version of the plugin or theme.
func GetVersionChangelog(raw string) []string {
	re := regexp.MustCompile(MatchVersionChangelog)

	return re.FindStringSubmatch(raw)
}

// GetVersionReleaseLog :: This function searches for the version of the plugin or theme.
func GetVersionReleaseLog(raw string) []string {
	re := regexp.MustCompile(MatchVersionReleaseLog)

	return re.FindStringSubmatch(raw)
}

// ReadAllFile :: This function will be responsible for reading the files.
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

func ReadCSVFile(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		printer.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		printer.Fatal(err)
	}

	return records
}
