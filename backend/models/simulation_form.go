package models

type FormField struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	Placeholder  string      `json:"placeholder"`
}

type FormStructure struct {
	Fields []FormField `json:"fields"`
}
