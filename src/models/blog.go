package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogItem struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:”title”`
	Content  string             `json:”content,omitempty”`
}
