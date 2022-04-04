package finders_users

import (
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/text"
)

func Run(target string) (*[]models.FindersModel) {
	var users = &[]models.FindersModel{}

	var json_api_models          = json_api(target)
	var rest_route_models        = rest_route(target)
	var rss_generator_models     = rss_generator(target)
	var yoast_seo_sitemap_models = yoast_seo_sitemap(target)

	var models_finders_all = [4](*[]models.FindersModel){json_api_models, rest_route_models, rss_generator_models, yoast_seo_sitemap_models}

	for _, models_finders := range models_finders_all {
		for _, models := range *models_finders {
			var exists_slug, _ = text.ContainsFindersSlug(*users, models.Slug)
			var exists_name, _ = text.ContainsFindersName(*users, models.Name)
	
			if !exists_slug || !exists_name {
				*users = append(*users, models)
			}
		}	
	}

	return users
}