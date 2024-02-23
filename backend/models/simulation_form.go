package models

type FormField struct {
    Name         string      `json:"name"`
    Type         string      `json:"type"`
    DefaultValue interface{} `json:"defaultValue"`
    Placeholder  string      `json:"placeholder"`
}

type FormCategory struct {
    CategoryName string      `json:"categoryName"`
    Fields       []FormField `json:"fields"`
}

type FormStructure struct {
    Categories []FormCategory `json:"categories"`
}
