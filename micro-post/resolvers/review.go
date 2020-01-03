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
	m models.Review
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
	ids  []graphql.ID
	from int
	to   int
	DB   *db.Services
}

func (r *Query) Review(ctx context.Context, args struct{ ID *string }) (*ReviewResolver, error) {
	id := *args.ID // dereference the pointer
	fmt.Println("id --", id)
	review, err := r.DB.GetReviewById(bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	s := ReviewResolver{
		m: *review,
		// Pass DB so we're able to use it in uderlying structs
		// DB: r.DB,
	}

	return &s, nil
}

// Reviews is the resolver for Reviews belonging to a Post (only getting the ids here)
func (p *PostResolver) Reviews(ctx context.Context) (*PostReviewsResolver, error) {
	ids := []graphql.ID{}

	reviews, _ := p.DB.GetReviewsByPostId(bson.M{"postId": "1"})

	for _, review := range *reviews {
		if review.PostID == p.m.ID {
			ids = append(ids, review.ID)
		}
	}

	s := PostReviewsResolver{
		ids: ids,
		// Pass DB so we're able to use it in uderlying structs
		DB: p.DB,
	}

	return &s, nil
}

// ReviewEdge is the edge containing all review nodes
type ReviewEdge struct {
	cursor graphql.ID
	node   ReviewResolver
}

func idInSlice(str graphql.ID, list []graphql.ID) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// Edges gives a list of all the review edges that belong to a post
func (u *PostReviewsResolver) Edges(ctx context.Context) (*[]*ReviewEdge, error) {
	selectedReviews := []*models.Review{}

	reviews, _ := u.DB.GetReviewsByPostId(bson.M{"postId": "1"})

	for _, review := range *reviews {
		if idInSlice(graphql.ID(review.ID), u.ids) {
			selectedReviews = append(selectedReviews, review)
		}
	}

	l := make([]*ReviewEdge, len(u.ids))
	for i := range l {
		l[i] = &ReviewEdge{
			cursor: helpers.EncodeCursor(i),
			node: ReviewResolver{
				m: *selectedReviews[i],
			},
		}
	}

	return &l, nil
}

// Cursor resolves the Cursor of a Review
func (u *ReviewEdge) Cursor(ctx context.Context) graphql.ID {
	return u.cursor
}

// Node resolves the Node of a Review
func (u *ReviewEdge) Node(ctx context.Context) *ReviewResolver {
	return &u.node
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
func (u *PageInfo) StartCursor(ctx context.Context) *graphql.ID {
	return &u.startCursor
}

// EndCursor ...
func (u *PageInfo) EndCursor(ctx context.Context) *graphql.ID {
	return &u.endCursor
}

// HasNextPage ...
func (u *PageInfo) HasNextPage(ctx context.Context) bool {
	return u.hasNextPage
}

// HasPreviousPage ...
func (u *PageInfo) HasPreviousPage(ctx context.Context) bool {
	return u.hasPreviousPage
}

// PageInfo resolves page info of a Review connection
func (u *PostReviewsResolver) PageInfo(ctx context.Context) (*PageInfo, error) {
	p := PageInfo{
		startCursor:     helpers.EncodeCursor(u.from),
		endCursor:       helpers.EncodeCursor(u.to - 1),
		hasNextPage:     u.to < len(u.ids),
		hasPreviousPage: u.from > 0,
	}
	return &p, nil
}
