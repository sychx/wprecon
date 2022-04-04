package finders_themes

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/printer"
	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/tools/interesting"
)

func indexof_theme_page(models_finders *[]models.FindersModel) *[]models.FindersModel {
	if directory_response, err := interesting.DirectoryThemes(); directory_response.Status == 200 {
		var regex = regexp.MustCompile("<a href=\"(.*?)/\">.*?/</a>")

		for _, regex_submatch := range regex.FindAllStringSubmatch(directory_response.Raw, -1) {
			if exists, _ := text.ContainsFindersName(*models_finders, regex_submatch[1]); !exists {
				*models_finders = append(*models_finders, models.FindersModel{
					Name: regex_submatch[1],
					FoundBy: "In the HTML - No Version",
					Others: []models.FindersOthersModel{ { Match: []string{} } },
				})
			}
		}

	} else if err != nil { printer.Danger(err) }

	return models_finders
}