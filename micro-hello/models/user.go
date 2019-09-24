package models

import "github.com/graph-gophers/graphql-go"

type User struct {
	ID       graphql.ID `json:"id"	bson:"id,omitempty"`
	LastName string     `json:"lastName"	bson:"lastName,omitempty"`
}
