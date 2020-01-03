package resolvers

import (
	"context"

	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type PostResolver struct {
	DB *db.Services
	m  models.Post
}

func (r *Query) Post(ctx context.Context, args struct{ ID *string }) (*PostResolver, error) {
	id := *args.ID // dereference the pointer

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

// Reviews is the resolver for Reviews belonging to a Post (only getting the ids here)
func (r *ReviewResolver) Post(ctx context.Context) (*PostResolver, error) {
	post, err := r.DB.GetPostById(bson.M{"id": "1"})
	spew.Dump("HEHEHEHEHEH", post)
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
