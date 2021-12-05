package config

import (
	_ "embed"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/printer"
	yaml "gopkg.in/yaml.v2"
)

//go:embed config.yaml
var config_yaml []byte

func GetConfig() *models.ConfigModel {
	var model models.ConfigModel

	var err = yaml.Unmarshal(config_yaml, &model)

	if err != nil {
		printer.Fatal(err)
	}

	return &model
}