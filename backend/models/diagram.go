package models

type DiagramNode struct {
	Key   string `json:"name"`
	Color string `json:"color"`
}

type DiagramLink struct {
	From string `json:"from"`
	To   string `json:"to"`
}
