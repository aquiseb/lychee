package resolvers

import (
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/astenmies/lychee/types"
	"github.com/graph-gophers/graphql-go"
)

//////////////////// NODE ////////////////////

// NodeResolver contains the node for resolving
type NodeResolver struct {
	node types.Node
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
