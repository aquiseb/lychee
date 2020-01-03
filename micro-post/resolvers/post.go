package resolvers

import (
	"context"
	"fmt"

	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type PostResolver struct {
	m  models.Post
	DB *db.Services
}

func (r *Query) Post(ctx context.Context, args struct{ ID *string }) (*PostResolver, error) {
	id := *args.ID // dereference the pointer
	fmt.Println("id --", id)
	post, err := r.DB.GetPostById(bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	s := PostResolver{
		m: *post,
		// Pass DB when the resolver below needs it as well!
		// For instance when { post(id: "1") { reviews { post { id } } } }
		DB: r.DB,
	}

	return &s, nil
}

// Title resolves the title field for Post
func (r *PostResolver) Title() *string {
	return &r.m.Title
}

// ID resolves the ID field for Post
func (r *PostResolver) ID() graphql.ID {
	return graphql.ID(r.m.ID)
}
