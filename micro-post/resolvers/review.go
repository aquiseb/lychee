package resolvers

import (
	"context"

	"github.com/astenmies/lychee/helpers"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

//////////////////// QUERY RELATIONSHIP ////////////////////

// Review resolver
func (q *Query) Review(ctx context.Context, args struct{ ID *string }) (*ReviewResolver, error) {
	id := *args.ID // dereference the pointer
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

//////////////////// POST RELATIONSHIP ////////////////////

// Reviews is the resolver for Reviews belonging to a Post
func (p *PostResolver) Reviews(ctx context.Context) (*PostReviewsResolver, error) {
	// ids := []graphql.ID{}

	reviews, _ := p.DB.GetReviewsByPostId(bson.M{"postId": p.m.ID})

	// for _, review := range *reviews {
	// 	if review.PostID == p.m.ID {
	// 		ids = append(ids, review.ID)
	// 	}
	// }

	s := PostReviewsResolver{
		// ids:     ids,
		DB:      p.DB,
		reviews: reviews,
	}

	return &s, nil
}

//////////////////// USER RELATIONSHIP ////////////////////

// Reviews is the resolver for Reviews belonging to a Post
func (p *UserResolver) Reviews(ctx context.Context) (*PostReviewsResolver, error) {
	spew.Dump("USER RESOLVER --", p.m.ID)
	reviews, _ := p.DB.GetReviewsByUserId(bson.M{"userId": p.m.ID})

	// [TODO] change the name of this resolver to something like ReviewConnectionResolver
	s := PostReviewsResolver{
		DB:      p.DB,
		reviews: reviews,
	}

	return &s, nil
}

//////////////////// REVIEW ////////////////////

// Edges gives a list of all the review edges that belong to a post
func (p *PostReviewsResolver) Edges(ctx context.Context) (*[]*ReviewEdge, error) {
	selectedReviews := []*models.Review{}
	reviews := p.reviews

	for _, review := range *reviews {
		selectedReviews = append(selectedReviews, review)
	}

	// [TODO] improve this
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
