package models

type DiagramNode struct {
	Key   string `json:"key" bson:"key"`
	Color string `json:"color" bson:"color"`
}

type DiagramLink struct {
	From string `json:"from"`
	To   string `json:"to"`
}
