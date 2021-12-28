package models

type EnumerateModel struct {
	Name       string
	FoundBy    string
	Others     []EnumerateOthersModel
}

type EnumerateOthersModel struct {
	Confidence int8
	Version    string
	FoundBy    string
	Match      []string	
}
