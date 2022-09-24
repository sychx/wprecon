package finders

var (
	Plugins = new([]Finders)
	Themes = new([]Finders)
)

type Finders struct {
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	FoundBy    string
	Others     []FindersOthers
}

type FindersOthers struct {
	Confidence int8
	Version    string
	FoundBy    string
	Match      []string	
}

type FindersVersion struct {
	Version string
	Match string
}