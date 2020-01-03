package resolvers

import (
	"github.com/graph-gophers/graphql-go"
)

func (u *UserResolver) ID() graphql.ID {
	return u.m.ID
}

func (u *UserResolver) LastName() string {
	return u.m.LastName
}
