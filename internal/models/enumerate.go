package models

type EnumerateModel struct {
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	FoundBy    string
	Others     []EnumerateOthersModel
}

type EnumerateOthersModel struct {
	Confidence int8
	Version    string
	FoundBy    string
	Match      []string	
}
