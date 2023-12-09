package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DiagramNode struct {
	ID    string  `json:"id" bson:"_id"`
	Color string  `json:"color" bson:"color"`
	Fx    float64 `json:"fx,omitempty" bson:"fx,omitempty"` // Fixed x-coordinate (optional)
	Fy    float64 `json:"fy,omitempty" bson:"fy,omitempty"` // Fixed y-coordinate (optional)
}

type DiagramLink struct {
	Source string `json:"source" bson:"source"`
	Target string `json:"target" bson:"target"`
}

type DiagramStructure struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Nodes []DiagramNode      `json:"nodes"`
	Links []DiagramLink      `json:"links"`
}
