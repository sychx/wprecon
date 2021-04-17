package extensions

import (
	"time"

	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/text"
	"github.com/blackbinn/wprecon/internal/pkg/wordlist"
)

// GetVersionByRequest :: It will search for versions of plugins or themes based on existing files. For example: readme, releaselog and etc...
// Revisar o funcionamento !
func GetVersionByRequest(target, path string) []string {
	if response := gohttp.SimpleRequest(target, path); response.Response.StatusCode == 200 && response.Raw != "" {
		if slice := text.GetVersionStableTag(response.Raw); len(slice) != 0 {
			return slice
		} else if slice := text.GetVersionChangelog(response.Raw); len(slice) != 0 {
			return slice
		} else if slice := text.GetVersionReleaseLog(response.Raw); len(slice) != 0 {
			return slice
		}
	}

	return []string{}
}

// GetVersionByChangeLogs :: It will search for plugin versions and themes based on existing files, making a file brute-force.
// Revisar o funcionamento !
func GetVersionByChangeLogs(target, path string) (string, string) {
	var channel = make(chan []string)

	for _, value := range wordlist.WPchangesLogs {
		go func() {
			if slice := GetVersionByRequest(target, path+value); len(slice) != 0 {
				channel <- slice
			}
		}()

		for {
			select {
			case i := <-channel:
				return i[0], i[1]
			case <-time.After(time.Second * 5):
				return "", ""
			}
		}
	}

	return "", ""
}

// GetVersionByReadme :: It will search for plugin versions and themes based on existing files, making a file brute-force.
// Revisar o funcionamento !
func GetVersionByReadme(target, path string) (string, string) {
	var channel = make(chan []string)

	for _, value := range wordlist.WPreadme {
		go func() {
			if slice := GetVersionByRequest(target, path+value); len(slice) != 0 {
				channel <- slice
			}
		}()

		for {
			select {
			case i := <-channel:
				return i[0], i[1]
			case <-time.After(time.Second * 5):
				return "", ""
			}
		}
	}

	return "", ""
}

// GetVersionByReleaseLog :: It will search for plugin versions and themes based on existing files, making a file brute-force.
// Revisar o funcionamento !
func GetVersionByReleaseLog(target, path string) (string, string) {
	var channel = make(chan []string)

	for _, value := range wordlist.WPreleaseLog {
		go func() {
			if slice := GetVersionByRequest(target, path+value); len(slice) != 0 {
				channel <- slice
			}
		}()

		for {
			select {
			case i := <-channel:
				return i[0], i[1]
			case <-time.After(time.Second * 5):
				return "", ""
			}
		}
	}

	return "", ""
}

// GetVersionByIndexOf ::
// Revisar o funcionamento !
func GetVersionByIndexOf(target, path string) (string, string) {
	var raw = gohttp.SimpleRequest(target, path).Raw

	if file := text.GetOneImportantFile(raw); file != "" {
		var raw = gohttp.SimpleRequest(target, path+file).Raw

		if slice := text.GetVersionChangelog(raw); len(slice) != 0 {
			return slice[0], slice[1]
		} else if slice := text.GetVersionStableTag(raw); len(slice) != 0 {
			return slice[0], slice[1]
		} else if slice := text.GetVersionChangelog(raw); len(slice) != 0 {
			return slice[0], slice[1]
		}
	}

	return "", ""
}
