package resolvers

import (
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
)

//////////////////// NODE ////////////////////

type Node interface {
	ID() graphql.ID
}

type NodeResolver struct {
	node Node
}

//////////////////// QUERY ////////////////////

type Query struct {
	DB *db.Services
}

//////////////////// REVIEW ////////////////////

// ReviewResolver resolves review
type ReviewResolver struct {
	DB *db.Services
	m  models.Review
}

// ReviewEdge is the edge containing all review nodes
type ReviewEdge struct {
	cursor graphql.ID
	node   ReviewResolver
}

//////////////////// POST ////////////////////

type PostResolver struct {
	DB *db.Services
	m  models.Post
}

// PostReviewsResolver gets all the Reviews that are connected to a certain Post
type PostReviewsResolver struct {
	DB      *db.Services
	reviews *[]*models.Review
	// ids     []graphql.ID
	from int
	to   int
}

//////////////////// USER ////////////////////

type UserResolver struct {
	DB *db.Services
	m  models.User
}

//////////////////// CONNECTIONS ////////////////////

// PageInfo provides information on our edge for pagination
type PageInfo struct {
	startCursor     graphql.ID
	endCursor       graphql.ID
	hasNextPage     bool
	hasPreviousPage bool
}
