package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Author string             `json:"author,omitempty"`
}
