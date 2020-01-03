package resolvers

import (
	"context"
	"fmt"

	"github.com/astenmies/lychee/helpers"
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type ReviewResolver struct {
	DB *db.Services
	m  models.Review
}

// var reviews = map[string]*models.Review{
// 	"1": {
// 		ID:     "11",
// 		Stars:  5,
// 		PostID: "2",
// 	},
// 	"2": {
// 		ID:     "22",
// 		Stars:  4,
// 		PostID: "1",
// 	},
// }

// PostReviewsResolver represents all the Reviews that are connected to a certain Post
type PostReviewsResolver struct {
	DB      *db.Services
	reviews *[]*models.Review
	// ids     []graphql.ID
	from int
	to   int
}

func (q *Query) Review(ctx context.Context, args struct{ ID *string }) (*ReviewResolver, error) {
	id := *args.ID // dereference the pointer
	fmt.Println("id --", id)
	review, err := q.DB.GetReviewById(bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	s := ReviewResolver{
		// Pass DB so we're able to use it in uderlying structs
		DB: q.DB,
		m:  *review,
	}

	return &s, nil
}

// Reviews is the resolver for Reviews belonging to a Post
func (p *PostResolver) Reviews(ctx context.Context) (*PostReviewsResolver, error) {
	ids := []graphql.ID{}

	reviews, _ := p.DB.GetReviewsByPostId(bson.M{"postId": p.m.ID})

	for _, review := range *reviews {
		if review.PostID == p.m.ID {
			ids = append(ids, review.ID)
		}
	}

	s := PostReviewsResolver{
		// ids:     ids,
		DB:      p.DB,
		reviews: reviews,
	}

	return &s, nil
}

// ReviewEdge is the edge containing all review nodes
type ReviewEdge struct {
	cursor graphql.ID
	node   ReviewResolver
}

// Edges gives a list of all the review edges that belong to a post
func (p *PostReviewsResolver) Edges(ctx context.Context) (*[]*ReviewEdge, error) {
	selectedReviews := []*models.Review{}
	reviews := p.reviews

	for _, review := range *reviews {
		selectedReviews = append(selectedReviews, review)
	}

	// [TODO] improve this. We don't use ids anymore, but `reviews` is directly passed to `Edges`
	l := make([]*ReviewEdge, len(*p.reviews))
	for i := range l {
		l[i] = &ReviewEdge{
			cursor: helpers.EncodeCursor(i),
			node: ReviewResolver{
				DB: p.DB,
				m:  *selectedReviews[i],
			},
		}
	}

	return &l, nil
}

// Cursor resolves the Cursor of a Review
func (r *ReviewEdge) Cursor(ctx context.Context) graphql.ID {
	return r.cursor
}

// Node resolves the Node of a Review
func (r *ReviewEdge) Node(ctx context.Context) *ReviewResolver {
	return &r.node
}

// ID resolves the ID of a Review
func (r *ReviewResolver) ID() graphql.ID {
	return graphql.ID(r.m.ID)
}

// Stars resolves the Stars of a review
func (r *ReviewResolver) Stars() *int32 {
	// [TODO] Edges should access r.m the same way
	stars := int32(r.m.Stars)
	return &stars
}

// PageInfo provides information on our edge for pagination
type PageInfo struct {
	startCursor     graphql.ID
	endCursor       graphql.ID
	hasNextPage     bool
	hasPreviousPage bool
}

// StartCursor ...
func (p *PageInfo) StartCursor(ctx context.Context) *graphql.ID {
	return &p.startCursor
}

// EndCursor ...
func (p *PageInfo) EndCursor(ctx context.Context) *graphql.ID {
	return &p.endCursor
}

// HasNextPage ...
func (p *PageInfo) HasNextPage(ctx context.Context) bool {
	return p.hasNextPage
}

// HasPreviousPage ...
func (p *PageInfo) HasPreviousPage(ctx context.Context) bool {
	return p.hasPreviousPage
}

// PageInfo resolves page info of a Review connection
func (p *PostReviewsResolver) PageInfo(ctx context.Context) (*PageInfo, error) {
	pi := PageInfo{
		startCursor:     helpers.EncodeCursor(p.from),
		endCursor:       helpers.EncodeCursor(p.to - 1),
		hasNextPage:     p.to < len(*p.reviews),
		hasPreviousPage: p.from > 0,
	}
	return &pi, nil
}
