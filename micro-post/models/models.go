package models

import "github.com/graph-gophers/graphql-go"

// Post ...
type Post struct {
	ID    graphql.ID `json:"id" bson:"id,omitempty"`
	Title string     `json:"title" bson:"title,omitempty"`
}

// User ...
type User struct {
	ID       graphql.ID `json:"id" bson:"id, omitempty"`
	LastName string     `json:"lastName" bson:"lastName,omitempty"`
}

// Review ...
type Review struct {
	ID     graphql.ID `json:"id" bson:"id, omitempty"`
	Stars  int        `json:"stars" bson:"stars,omitempty"`
	PostID graphql.ID `json:"postId" bson:"postId,omitempty"`
}
