package finders_users

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
)

func rss_generator(target string) (models_finders *[]models.FindersModel) {
	models_finders = &[]models.FindersModel{}

	if response := net.SimpleRequest(target+"/feed"); response.Response.StatusCode == 200 && response.Raw != "" {
		var regex = regexp.MustCompile("<dc:creator><!\\[CDATA\\[(.+?)\\]\\]></dc:creator>")

		for _, match := range regex.FindAllStringSubmatch(response.Raw, -1) {
			*models_finders = append(*models_finders, models.FindersModel{ Name: match[1], FoundBy: "RSS Generator" })
		}
	}

	return
}