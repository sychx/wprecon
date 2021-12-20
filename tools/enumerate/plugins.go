package enumerate

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/tools/findings"
	"github.com/blackcrw/wprecon/tools/interesting"
)

/*
	plugin[0] :: match
	plugin[1] :: name
	plugin[2] :: version
*/

func PluginPassive() *[]models.EnumerateModel {
	var model []models.EnumerateModel

	var regxp = regexp.MustCompile(`wp-content/plugins/(.*?)/.*?[css|js].*?ver=(\d{1,2}\.\d{1,2}\.\d{1,3})`)
	
	for _, plugin_submatch := range regxp.FindAllStringSubmatch(database.Memory.GetString("HTTP Index Raw"), -1) {
		var matriz models.EnumerateModel
		var matriz_others models.EnumerateOthersModel

		if has_name, int_name := text.ContainsEnumerateName(model, plugin_submatch[1]); has_name {
			var has_version, _ = text.ContainsEnumerateVersion(model, plugin_submatch[2])
			
			if !has_version {
				matriz_others.Version = plugin_submatch[2]
				matriz_others.Confidence += 10
				matriz_others.Match = append(matriz_others.Match, plugin_submatch[0])
				
				model[int_name].Others = append(model[int_name].Others, matriz_others)
			}

			var has_match, int_match = text.ContainsEnumerateMatch(model, plugin_submatch[0])

			if !has_match && model[int_name].Others[int_match].Version == plugin_submatch[2] {
				model[int_name].Others[int_match].Match = append(model[int_name].Others[int_match].Match, plugin_submatch[0])
				model[int_name].Others[int_match].Confidence += 10
			}
		} else {
			matriz_others.Confidence += 10
			matriz.FoundBy = "In the HTML of the index"

			matriz_others.Version = plugin_submatch[2]
			matriz_others.Match = append(matriz_others.Match, plugin_submatch[0])
			
			matriz.Name = plugin_submatch[1]
			matriz.Others = append(matriz.Others, matriz_others)

			model = append(model, matriz)
		}
	}

	return &model
}

func PluginAggressive() *[]models.EnumerateModel {
	var model []models.EnumerateModel

	model = *PluginPassive() 

	if directory_response := interesting.DirectoryPlugins(); directory_response.Status == 200 {
		var regxp = regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")
		var matriz models.EnumerateModel

		for _, plugin_submatch := range regxp.FindAllStringSubmatch(directory_response.Raw, -1) {
			if has, _ := text.ContainsEnumerateName(model, plugin_submatch[1]); !has {
				matriz.Name = plugin_submatch[1]
				matriz.FoundBy = "In the HTML of the index - No version"
				model = append(model, matriz)
			}
		}
	}

	var regxp = regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/plugins/(.*?)/.*?[.css|.js]")

	for _, plugin_submatch := range regxp.FindAllStringSubmatch(database.Memory.GetString("HTTP Index Raw"), -1) {
		var matriz models.EnumerateModel

		if has_name,_ := text.ContainsEnumerateName(model, plugin_submatch[1]); !has_name {
			matriz.Name = plugin_submatch[1]

			model = append(model, matriz)
		}
	}

	wg.Add(len(model))

	for _, model_plugins := range model {
		go func(model_plugins models.EnumerateModel){
			var path = database.Memory.GetString("HTTP wp-content") + "/plugins/" + model_plugins.Name + "/"
		
			var model_finding_changeslogs = findings.FindingVersionByChangesLogs(path)
			var model_finding_releaselog = findings.FindingVersionByReleaseLog(path)
			var model_finding_indexof = findings.FindingVersionByIndexOf(path)
			var model_finding_readme = findings.FindingVersionByReadme(path)
					
			findings_add_version(model_plugins.Name, &model, model_finding_changeslogs)
			findings_add_version(model_plugins.Name, &model, model_finding_releaselog)
			findings_add_version(model_plugins.Name, &model, model_finding_indexof)
			findings_add_version(model_plugins.Name, &model, model_finding_readme)
		
			defer wg.Done()
		}(model_plugins)
	}

	wg.Wait()

	return &model
}