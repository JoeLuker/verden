package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DiagramNode struct {
	Key   string `json:"key" bson:"key"`
	Color string `json:"color" bson:"color"`
}

type DiagramLink struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type DiagramStructure struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Nodes []DiagramNode      `json:"nodes"`
	Links []DiagramLink      `json:"links"`
}
