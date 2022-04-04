package finders_users

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
)

func yoast_seo_sitemap(target string) (models_finders *[]models.FindersModel) {
	models_finders = &[]models.FindersModel{}

	var target_host = net.GetHost(target)
	var response = net.SimpleRequest(target+"/wp-content/cache/page_enhanced/"+target_host+"/author-sitemap.xml/_index_ssl.xml")

	if response.Response.StatusCode == 200 {
		var regex = regexp.MustCompile("<loc>https?://.*?/author/(.*?)/</loc>")

		for _, regex_submatch := range regex.FindAllStringSubmatch(response.Raw, -1) {
			*models_finders = append(*models_finders, models.FindersModel{Slug: regex_submatch[1], FoundBy: "Yoast SEO Sitemap (Cached)"})
		}
	}

	return
}