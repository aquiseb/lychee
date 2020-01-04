package models

import "github.com/graph-gophers/graphql-go"

type User struct {
	ID        graphql.ID `json:"id" bson:"id,omitempty"`
	Firstname string     `json:"firstname" bson:"firstname,omitempty"`
	Lastname  string     `json:"lastname" bson:"lastname,omitempty"`
}
