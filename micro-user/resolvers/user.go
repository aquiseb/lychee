package resolvers

import (
	"context"

	"github.com/aquiseb/lychee/micro-user/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

var us = map[string]*models.User{
	"1": {
		ID:        "1",
		Firstname: "Aline",
		Lastname:  "Hoho",
	},
}

// Edges gives a list of all the review edges that belong to a post
// [TODO] NEEDS MASSIVE CLEANUP AND micro-post Edges should maybe have the same approach
func (p *Query) AllUsers(ctx context.Context) ([]*UserResolver, error) {
	selectedReviews := []*UserResolver{}
	users, _ := p.DB.GetAllUsers()

	for _, user := range *users {
		selectedReviews = append(selectedReviews, &UserResolver{m: *user})
	}

	return selectedReviews, nil
}

// User resolves the post query
func (q *Query) User(ctx context.Context, args struct{ ID *string }) (*UserResolver, error) {
	id := *args.ID // dereferences the pointer

	user, err := q.DB.GetUserByID(bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	s := UserResolver{
		// Pass DB when the resolver below needs it as well!
		// For instance when { user(id: "1") { reviews { post { id } } } }
		DB: q.DB,
		m:  *user,
	}

	return &s, nil
}

func (u *UserResolver) ID() graphql.ID {
	return u.m.ID
}

func (u *UserResolver) Firstname() string {
	return u.m.Firstname
}

func (u *UserResolver) Lastname() string {
	return u.m.Lastname
}
