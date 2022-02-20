package enumerate

import (
	"encoding/json"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/text"
)

func UserAggressive() *[]models.EnumerateModel {
	var model_enum []models.EnumerateModel
	var model_json []models.EnumerateModel
	var target = database.Memory.GetString("Options URL")

	var response = net.SimpleRequest(target+"wp-json/wp/v2/users")

	if response.Response.StatusCode == 200 && response.Raw != "" {
		json.Unmarshal([]byte(response.Raw), &model_enum)
	}

	response = net.SimpleRequest(target+"?rest_route=/wp/v2/users")
	
	if response.Response.StatusCode == 200 && response.Raw != "" {
		json.Unmarshal([]byte(response.Raw), &model_json)

		for _, model := range model_json {
			if has_user, _ := text.ContainsEnumerateSlug(model_enum, model.Slug); !has_user {
				model_enum = append_for(model_enum, model_json)
			}
		}		
	}

	return &model_enum
}
