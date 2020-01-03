package resolvers

import (
	"context"

	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type PostResolver struct {
	DB *db.Services
	m  models.Post
}

func (q *Query) Post(ctx context.Context, args struct{ ID *string }) (*PostResolver, error) {
	id := *args.ID // dereference the pointer

	post, err := q.DB.GetPostById(bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	s := PostResolver{
		// Pass DB when the resolver below needs it as well!
		// For instance when { post(id: "1") { reviews { post { id } } } }
		DB: q.DB,
		m:  *post,
	}

	return &s, nil
}

// Post is the resolver for Post belonging to a Review
func (r *ReviewResolver) Post(ctx context.Context) (*PostResolver, error) {
	post, err := r.DB.GetPostById(bson.M{"id": r.m.ID})
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
func (p *PostResolver) Title() *string {
	return &p.m.Title
}

// ID resolves the ID field for Post
func (p *PostResolver) ID() graphql.ID {
	return graphql.ID(p.m.ID)
}
