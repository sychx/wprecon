package models

type FindersModel struct {
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	FoundBy    string
	Others     []FindersOthersModel
}

type FindersOthersModel struct {
	Confidence int8
	Version    string
	FoundBy    string
	Match      []string	
}

type FindersVersionModel struct {
	Version string
	Match string
}