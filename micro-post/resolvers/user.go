package resolvers

import (
	"github.com/graph-gophers/graphql-go"
)

func (u *UserResolver) ID() graphql.ID {
	return u.m.ID
}

func (u *UserResolver) Lastname() string {
	return u.m.Lastname
}
