package models

type MiddlewareFirewallModel struct {
	Name       string
	Solve      string
	Confidence int
	FoundBy    string
}

type InterestingModel struct {
	Name       string
	Url        string
	Status     int
	Raw        string
	Confidence int8
	FoundBy    string
}

type FindingsVersionModel struct {
	Version string
	Match string
}

type ConfigModel struct {
	App struct {
		Name        string `yaml:"name"`
		Author      string `yaml:"author"`
		Description string `yaml:"description"`
		Version     string `yaml:"version"`
		ApiUrl      string `yaml:"api_url"`
	} `yaml:"application"`
}