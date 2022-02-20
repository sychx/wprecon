package enumerate

import (
	"fmt"
	"regexp"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/printer"
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
	var model_enum []models.EnumerateModel

	var regxp = regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + `/plugins/(.*?)/.*?[css|js].*?ver=(\d{1,2}\.\d{1,2}\.\d{1,3})`)
	
	for _, plugin_submatch := range regxp.FindAllStringSubmatch(database.Memory.GetString("HTTP Index Raw"), -1) {
		var model_matriz models.EnumerateModel
		var model_matriz_others models.EnumerateOthersModel
	
		if has_name, int_name := text.ContainsEnumerateName(model_enum, plugin_submatch[1]); has_name {
			var has_version, _ = text.ContainsEnumerateVersion(model_enum, plugin_submatch[2])
			
			if !has_version {
				model_matriz_others.Version = plugin_submatch[2]
				model_matriz_others.FoundBy = "Version in the HTML code of the index"
				model_matriz_others.Match = append(model_matriz_others.Match, plugin_submatch[0])
				model_matriz_others.Confidence = format_confidence(model_matriz_others.Confidence, 10)
				
				model_enum[int_name].Others = append(model_enum[int_name].Others, model_matriz_others)
			}

			var has_match, int_match = text.ContainsEnumerateMatch(model_enum, plugin_submatch[0])

			if !has_match && model_enum[int_name].Others[int_match].Version == plugin_submatch[2] {
				model_enum[int_name].Others[int_match].Match = append(model_enum[int_name].Others[int_match].Match, plugin_submatch[0])
				model_enum[int_name].Others[int_match].Confidence = format_confidence(model_enum[int_name].Others[int_match].Confidence, 10)
			}
		} else {
			model_matriz.FoundBy = "In the HTML of the index"
			model_matriz_others.Confidence = format_confidence(model_matriz_others.Confidence, 10)

			model_matriz_others.Version = plugin_submatch[2]
			model_matriz_others.Match = append(model_matriz_others.Match, plugin_submatch[0])
			
			model_matriz.Name = plugin_submatch[1]
			model_matriz.Others = append(model_matriz.Others, model_matriz_others)

			model_enum = append(model_enum, model_matriz)
		}
	}

	regxp = regexp.MustCompile(database.Memory.GetString("HTTP wp-content") + "/plugins/(.*?)/.*?[.css|.js]")
	
	for _, plugin_submatch := range regxp.FindAllStringSubmatch(database.Memory.GetString("HTTP Index Raw"), -1) {
		var model_matriz models.EnumerateModel
	
		if has_name,_ := text.ContainsEnumerateName(model_enum, plugin_submatch[1]); !has_name {
			model_matriz.Name = plugin_submatch[1]
			model_matriz.FoundBy = "In the HTML of the index - No version"
			model_enum = append(model_enum, model_matriz)
		}
	}

	return &model_enum
}

func PluginAggressive() *[]models.EnumerateModel {
	var model_enum []models.EnumerateModel = *PluginPassive()

	if directory_response, err := interesting.DirectoryPlugins(); directory_response.Status == 200 {
		var regxp = regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")
		var model_matriz models.EnumerateModel

		for _, plugin_submatch := range regxp.FindAllStringSubmatch(directory_response.Raw, -1) {
			if has, _ := text.ContainsEnumerateName(model_enum, plugin_submatch[1]); !has {
				model_matriz.Name = plugin_submatch[1]
				model_matriz.FoundBy = "In the HTML - No version"
				model_enum = append(model_enum, model_matriz)
			}
		}
	} else if err != nil { printer.Danger(fmt.Sprintf("%s", err)) }

	wg.Add(len(model_enum))

	for _, model_plugins := range model_enum {
		go func(model_plugins models.EnumerateModel){
			var path = database.Memory.GetString("HTTP wp-content") + "/plugins/" + model_plugins.Name + "/"
		
			var model_finding_changeslogs = findings.FindingVersionByChangesLogs(path)
			var model_finding_releaselog = findings.FindingVersionByReleaseLog(path)
			var model_finding_indexof = findings.FindingVersionByIndexOf(path)
			var model_finding_readme = findings.FindingVersionByReadme(path)
					
			findings_add_version(model_plugins.Name, &model_enum, model_finding_changeslogs)
			findings_add_version(model_plugins.Name, &model_enum, model_finding_releaselog)
			findings_add_version(model_plugins.Name, &model_enum, model_finding_indexof)
			findings_add_version(model_plugins.Name, &model_enum, model_finding_readme)
		
			defer wg.Done()
		}(model_plugins)
	}

	wg.Wait()
	
	return &model_enum
}