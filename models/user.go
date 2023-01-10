package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty"`
	LastName string             `json:"lastname,omitempty"`
} //can't send data empty thanks to omitempty
