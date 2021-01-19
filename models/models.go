package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Book :Create Struct
type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"Title,omitempty" bson:"_id,omitempty"`
	Author *Author            `json:"Author,omitempty" bson:"_id,omitempty"`
}

//Author : Create Struct
type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
