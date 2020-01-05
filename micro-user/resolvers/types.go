package resolvers

import (
	"github.com/astenmies/lychee/micro-user/db"
	"github.com/astenmies/lychee/micro-user/models"
	"github.com/graph-gophers/graphql-go"
)

//////////////////// QUERY ////////////////////

// Query contains the DB so it's made available in all resolvers
type Query struct {
	DB *db.Services
}

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

// // NodeResolver contains the node for resolving
// type NodeResolver struct {
// 	types.Node
// }

// UserResolver contains the DB and the model for resolving
type UserResolver struct {
	DB *db.Services
	m  models.User
}

type UsersResolver struct {
	DB    *db.Services
	users *[]*models.User
	// ids     []graphql.ID
	from int
	to   int
}
