package finders_version

import (
	"sync"
	"time"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/internal/wordlist"
)

func GoByRequest(channel chan []string, path string) {
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

func ByReadme(path string) *models.FindersVersionModel {
	var channel = make(chan []string)
	
	for _, word := range wordlist.WPReadme {
		go GoByRequest(channel, path+word)

		select {
		case response := <-channel:
			return &models.FindersVersionModel{ Version: response[1], Match: response[0] }
		case <-time.After(time.Second*8):
			continue
		}
	}

	return &models.FindersVersionModel{}
}

func ByReleaseLog(path string) *models.FindersVersionModel {
	var channel = make(chan []string)
	
	for _, word := range wordlist.WPReleaseLog {
		go GoByRequest(channel, path+word)

		select {
		case response := <-channel:
			return &models.FindersVersionModel{ Version: response[1], Match: response[0] }
		case <-time.After(time.Second*8):
			continue
		}
	}

	return &models.FindersVersionModel{}
}

func ByChangesLogs(path string) *models.FindersVersionModel {
	var channel = make(chan []string)
	
	for _, word := range wordlist.WPChangesLogs {
		go GoByRequest(channel, path+word)

		select {
		case response := <-channel:
			return &models.FindersVersionModel{ Version: response[1], Match: response[0] }
		case <-time.After(time.Second*8):
			continue
		}
	}

	return &models.FindersVersionModel{}
}

func ByIndexOf(path string) *models.FindersVersionModel {
	var target = database.Memory.GetString("Options URL")
	var files = text.FindImportantFiles(net.SimpleRequest(target+path).Raw)

	for _, name_file := range files {
		var response_raw = net.SimpleRequest(target+path+name_file[0]).Raw
	
		if slice := text.GetVersionByChangelog(response_raw); len(slice) != 0 {
			return &models.FindersVersionModel{Version: slice[1], Match: slice[0]}
		} else if slice := text.GetVersionByStableTag(response_raw); len(slice) != 0 {
			return &models.FindersVersionModel{Version: slice[1], Match: slice[0]}
		} else if slice := text.GetVersionByChangelog(response_raw); len(slice) != 0 {
			return &models.FindersVersionModel{Version: slice[1], Match: slice[0]}
		}
	}

	return &models.FindersVersionModel{}
}

func Run(models_finders *[]models.FindersModel) *[]models.FindersModel {
	var sync_wait_group sync.WaitGroup

	var func_format_version = func(name string, models_finders *[]models.FindersModel, models_finders_version *models.FindersVersionModel) {
		if models_finders_version.Version != "" {
			var _, int_name = text.ContainsFindersName(*models_finders, name)

			if exists_version, int_version := text.ContainsFindersVersion(*models_finders, models_finders_version.Version); exists_version {
				(*models_finders)[int_name].Others[int_version].Confidence = text.FormatConfidence((*models_finders)[int_name].Others[int_version].Confidence, 20)
			} else {
				(*models_finders)[int_name].Others = append((*models_finders)[int_name].Others, models.FindersOthersModel{
					Version: models_finders_version.Version,
					FoundBy: "Findings Aggressive In Important Files",
					Match: append((*models_finders)[int_name].Others[int_version].Match, models_finders_version.Match),
					Confidence: text.FormatConfidence((*models_finders)[int_name].Others[int_version].Confidence, 40),
				})

			}
		}
	}

	sync_wait_group.Add(len(*models_finders))

	for _, model := range *models_finders {
		go func(model models.FindersModel){
			var path = database.Memory.GetString("HTTP wp-content") + "/plugins/" + model.Name + "/"
	
			var model_finders_changeslogs = ByChangesLogs(path)
			var model_finders_releaselog = ByReleaseLog(path)
			var model_finders_indexof = ByIndexOf(path)
			var model_finders_readme = ByReadme(path)

			func_format_version(model.Name, models_finders, model_finders_changeslogs)
			func_format_version(model.Name, models_finders, model_finders_releaselog)
			func_format_version(model.Name, models_finders, model_finders_indexof)
			func_format_version(model.Name, models_finders, model_finders_readme)
			
			defer sync_wait_group.Done()
		}(model)
	}

	sync_wait_group.Wait()

	return models_finders
}