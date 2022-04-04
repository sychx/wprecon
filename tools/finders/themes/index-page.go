package finders_themes

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/text"
)

func index_body(path_wp_content, index_raw string) (models_finders *[]models.FindersModel)	{
	models_finders = &[]models.FindersModel{}

	var regex = regexp.MustCompile(path_wp_content + "/themes/(.*?)/.*?[.css|.js]")

	for _, regex_submatch := range regex.FindAllStringSubmatch(index_raw, -1) {
		if exists_name, _ := text.ContainsFindersName(*models_finders, regex_submatch[1]); !exists_name {		
			*models_finders = append(*models_finders, models.FindersModel{
				Name: regex_submatch[1],
				FoundBy: "In the HTML of the index - No Version",
				Others: []models.FindersOthersModel{ { Match: []string{} } },
			})
		}
	}

	return
}

/* This function is a complement of the function: index_body(); It serves to enumerate how theme versions. */
func index_body_version(models_finders *[]models.FindersModel, path_wp_content, index_raw string) *[]models.FindersModel {
	var regex = regexp.MustCompile(path_wp_content + `/themes/(.*?)/.*?[css|js].*?ver=(\d{1,2}\.\d{1,2}\.\d{1,3})`)

	for _, regex_submatch := range regex.FindAllStringSubmatch(index_raw, -1) {
		if exists_name, int_name := text.ContainsFindersName(*models_finders, regex_submatch[1]); exists_name {
			var exists_match, int_match = text.ContainsFindersMatch(*models_finders, regex_submatch[0])

			if exists_version, _ := text.ContainsFindersVersion(*models_finders, regex_submatch[2]); !exists_version {
				(*models_finders)[int_name].FoundBy = "In the HTML of the index"

				(*models_finders)[int_name].Others = append((*models_finders)[int_name].Others, models.FindersOthersModel{
					Version: regex_submatch[2],
					FoundBy: "Version in the HTML code of the index",
					Match: append((*models_finders)[int_name].Others[int_match].Match, regex_submatch[0]),
					Confidence: text.FormatConfidence((*models_finders)[int_name].Others[int_match].Confidence, 15),
				})
			}

			if !exists_match && (*models_finders)[int_name].Others[int_match].Version == regex_submatch[2] {
				(*models_finders)[int_name].Others[int_match].Match = append((*models_finders)[int_name].Others[int_match].Match, regex_submatch[0])
				(*models_finders)[int_name].Others[int_match].Confidence = text.FormatConfidence((*models_finders)[int_name].Others[int_match].Confidence, 15)
			}
		} else {
			var _, int_match = text.ContainsFindersMatch(*models_finders, regex_submatch[0])

			*models_finders = append(*models_finders, models.FindersModel{
				Name: regex_submatch[1],
				FoundBy: "In the HTML of the index",
				Others: []models.FindersOthersModel{
					{
						Version: regex_submatch[2],
						FoundBy:"Version in the HTML code of the index",
						Match: append((*models_finders)[int_name].Others[int_match].Match, regex_submatch[0]),
						Confidence: text.FormatConfidence(0, 15),
					},
				},
			})
		}
	}

	return models_finders
}