package findings

import (
	"time"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/internal/wordlist"
)

func GoFindingVersionByRequest(channel chan []string, path string) {
	var response = net.SimpleRequest(database.Memory.GetString("Options URL")+path)

	if response.Response.StatusCode == 200 && response.Raw != "" {
		if slice := text.GetVersionByStableTag(response.Raw); len(slice) != 0 {
			channel <- slice
		} else if slice := text.GetVersionByChangelog(response.Raw); len(slice) != 0 {
			channel <- slice
		} else if slice := text.GetVersionByReleaseLog(response.Raw); len(slice) != 0 {
			channel <- slice
		}
	}
}

func FindingVersionByReadme(path string) *models.FindingsVersionModel {
	var channel = make(chan []string)
	
	for _, word := range wordlist.WPReadme {
		go GoFindingVersionByRequest(channel, path+word)

		select {
		case response := <-channel:
			return &models.FindingsVersionModel{ Version: response[1], Match: response[0] }
		case <-time.After(time.Second*8):
			continue
		}
	}

	return &models.FindingsVersionModel{}
}

func FindingVersionByReleaseLog(path string) *models.FindingsVersionModel {
	var channel = make(chan []string)
	
	for _, word := range wordlist.WPReleaseLog {
		go GoFindingVersionByRequest(channel, path+word)

		select {
		case response := <-channel:
			return &models.FindingsVersionModel{ Version: response[1], Match: response[0] }
		case <-time.After(time.Second*8):
			continue
		}
	}

	return &models.FindingsVersionModel{}
}

func FindingVersionByChangesLogs(path string) *models.FindingsVersionModel {
	var channel = make(chan []string)
	
	for _, word := range wordlist.WPChangesLogs {
		go GoFindingVersionByRequest(channel, path+word)

		select {
		case response := <-channel:
			return &models.FindingsVersionModel{ Version: response[1], Match: response[0] }
		case <-time.After(time.Second*8):
			continue
		}
	}

	return &models.FindingsVersionModel{}
}

func FindingVersionByIndexOf(path string) *models.FindingsVersionModel {
	var target = database.Memory.GetString("Options URL")
	var files = text.FindImportantFiles(net.SimpleRequest(target+path).Raw)

	for _, name_file := range files {
		var response_raw = net.SimpleRequest(target+path+name_file[0]).Raw
	
		if slice := text.GetVersionByChangelog(response_raw); len(slice) != 0 {
			return &models.FindingsVersionModel{Version: slice[1], Match: slice[0]}
		} else if slice := text.GetVersionByStableTag(response_raw); len(slice) != 0 {
			return &models.FindingsVersionModel{Version: slice[1], Match: slice[0]}
		} else if slice := text.GetVersionByChangelog(response_raw); len(slice) != 0 {
			return &models.FindingsVersionModel{Version: slice[1], Match: slice[0]}
		}
	}

	return &models.FindingsVersionModel{}
}
