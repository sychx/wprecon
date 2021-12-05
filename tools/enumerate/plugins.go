package enumerate

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/text"
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