package finders_themes

import (
	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	finders_versions "github.com/blackcrw/wprecon/tools/finders/versions"
)

func Passive(target string) *[]models.FindersModel {
	var index_raw = database.Memory.GetString("HTTP Index Raw")
	var path_wp_content = database.Memory.GetString("HTTP wp-content")

	return index_body_version(index_body(path_wp_content, index_raw), path_wp_content, index_raw)
}

func Aggressive(target string) *[]models.FindersModel {
	var index_raw = database.Memory.GetString("HTTP Index Raw")
	var path_wp_content = database.Memory.GetString("HTTP wp-content")

	return finders_versions.Run(indexof_theme_page(index_body_version(index_body(path_wp_content, index_raw), path_wp_content, index_raw)))
}