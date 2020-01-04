package resolvers

import (
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
)

//////////////////// NODE ////////////////////

// Node interface is required by nautilus gateway
// See here The Node interface here: https://itnext.io/a-guide-to-graphql-schema-federation-part-1-995b639ac035?
// The only requirement that schema federation imposes on a service is that it must satisfy the Relay Global Object Identification Specification.
// See here: https://facebook.github.io/relay/graphql/objectidentification.htm
type Node interface {
	ID() graphql.ID
}

// NodeResolver contains the node for resolving
type NodeResolver struct {
	node Node
}

//////////////////// QUERY ////////////////////

// Query contains the DB so it's made available in all resolvers
type Query struct {
	DB *db.Services
}

//////////////////// REVIEW ////////////////////

// ReviewResolver contains the DB and the model for resolving
type ReviewResolver struct {
	DB *db.Services
	m  models.Review
}

// ReviewEdge is the edge containing review nodes
type ReviewEdge struct {
	cursor graphql.ID
	node   ReviewResolver
}

//////////////////// POST ////////////////////

// PostResolver contains the DB and the model for resolving
type PostResolver struct {
	DB *db.Services
	m  models.Post
}

// PostReviewsResolver contains the DB and the models for resolving the reviews of a post
type PostReviewsResolver struct {
	DB      *db.Services
	reviews *[]*models.Review
	// ids     []graphql.ID
	from int
	to   int
}

// UserReviewsResolver contains the DB and the models for resolving the reviews of a user
type UserReviewsResolver struct {
	DB      *db.Services
	reviews *[]*models.Review
	// ids     []graphql.ID
	from int
	to   int
}

//////////////////// USER ////////////////////

// UserResolver contains the DB and the model for resolving
type UserResolver struct {
	DB *db.Services
	m  models.User
}

//////////////////// CONNECTIONS ////////////////////

// PageInfo provides information to the edge for pagination
type PageInfo struct {
	startCursor     graphql.ID
	endCursor       graphql.ID
	hasNextPage     bool
	hasPreviousPage bool
}
