package extensions

import (
	"time"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/text"
	"github.com/blackbinn/wprecon/pkg/wordlist"
)

// GetVersionByRequest ::
func GetVersionByRequest(path string) []string {
	if response := gohttp.SimpleRequest(database.Memory.GetString("Target"), path); response.Response.StatusCode == 200 && response.Raw != "" {
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

// GetVersionByChangeLogs ::
func GetVersionByChangeLogs(path string) (string, string) {
	channel := make(chan []string)

	for _, value := range wordlist.WPchangesLogs {
		go func() {
			if slice := GetVersionByRequest(path + value); len(slice) != 0 {
				channel <- slice
			}
		}()

		time.Sleep(time.Millisecond * 150)

		select {
		case i := <-channel:
			return i[0], i[1]
		default:
			return "", ""
		}
	}

	return "", ""
}

// GetVersionByReadme ::
func GetVersionByReadme(path string) (string, string) {
	channel := make(chan []string)

	for _, value := range wordlist.WPreadme {
		go func() {
			if slice := GetVersionByRequest(path + value); len(slice) != 0 {
				channel <- slice
			}
		}()

		time.Sleep(time.Millisecond * 150)

		select {
		case i := <-channel:
			return i[0], i[1]
		default:
			return "", ""
		}
	}

	return "", ""
}

// GetVersionByReleaseLog ::
func GetVersionByReleaseLog(path string) (string, string) {
	channel := make(chan []string)

	for _, value := range wordlist.WPreleaseLog {
		go func() {
			if slice := GetVersionByRequest(path + value); len(slice) != 0 {
				channel <- slice
			}
		}()

		time.Sleep(time.Millisecond * 150)

		select {
		case i := <-channel:
			return i[0], i[1]
		default:
			return "", ""
		}
	}

	return "", ""
}

// GetVersionByIndexOf ::
func GetVersionByIndexOf(path string) (string, string) {
	raw := gohttp.SimpleRequest(database.Memory.GetString("Target"), path).Raw

	if file := text.GetOneImportantFile(raw); file != "" {
		raw := gohttp.SimpleRequest(database.Memory.GetString("Target"), path+file).Raw

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
