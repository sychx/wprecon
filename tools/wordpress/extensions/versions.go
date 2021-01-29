package extensions

import (
	"time"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/text"
	"github.com/blackbinn/wprecon/pkg/wordlist"
)

// GetVersionByRequest ::
func GetVersionByRequest(path string) string {
	if response := gohttp.SimpleRequest(InfosWprecon.Target, path); response.Response.StatusCode == 200 && response.Raw != "" {
		if version := text.GetVersionStableTag(response.Raw); version != "" {
			return version
		} else if version := text.GetVersionChangelog(response.Raw); version != "" {
			return version
		} else if version := text.GetVersionReleaseLog(response.Raw); version != "" {
			return version
		}
	}

	return ""
}

// GetVersionByChangeLogs ::
func GetVersionByChangeLogs(path string) string {
	channel := make(chan string)

	for _, value := range wordlist.WPchangesLogs {
		go func() {
			if version := GetVersionByRequest(path + value); version != "" {
				channel <- version
			}
		}()

		time.Sleep(time.Millisecond * 100)

		select {
		case version := <-channel:
			return version
		default:
			return ""
		}
	}

	return ""
}

// GetVersionByReadme ::
func GetVersionByReadme(path string) string {
	channel := make(chan string)

	for _, value := range wordlist.WPreadme {
		go func() {
			if version := GetVersionByRequest(path + value); version != "" {
				channel <- version
			}
		}()

		time.Sleep(time.Millisecond * 100)

		select {
		case version := <-channel:
			return version
		default:
			return ""
		}
	}

	return ""
}

// GetVersionByReleaseLog ::
func GetVersionByReleaseLog(path string) string {
	channel := make(chan string)

	for _, value := range wordlist.WPreleaseLog {
		go func() {
			if version := GetVersionByRequest(path + value); version != "" {
				channel <- version
			}
		}()

		time.Sleep(time.Millisecond * 100)

		select {
		case version := <-channel:
			return version
		default:

			return ""
		}
	}

	return ""
}

// GetVersionByIndexOf ::
func GetVersionByIndexOf(path string) string {
	raw := gohttp.SimpleRequest(InfosWprecon.Target, path).Raw

	if file := text.GetOneImportantFile(raw); file != "" {
		raw := gohttp.SimpleRequest(InfosWprecon.Target, path+file).Raw

		if version := text.GetVersionChangelog(raw); version != "" {
			return version
		} else if version := text.GetVersionStableTag(raw); version != "" {
			return version
		} else if version := text.GetVersionChangelog(raw); version != "" {
			return version
		}
	}

	return ""
}
