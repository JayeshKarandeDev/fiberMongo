package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cars struct {
	Id    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Color string             `json:"color"`
}
