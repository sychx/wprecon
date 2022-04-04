package finders_users

import (
	"encoding/json"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
)

func json_api(target string) (models_finders *[]models.FindersModel) {
	var response = net.SimpleRequest(target+"/wp-json/wp/v2/users")

	if response.Response.StatusCode == 200 && response.Raw != "" {
		json.Unmarshal([]byte(response.Raw), &models_finders)
	}

	return &[]models.FindersModel{}
}