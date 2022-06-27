package resolvers

import (
	"github.com/aquiseb/lychee/micro-user/db"
	"github.com/aquiseb/lychee/micro-user/models"
	"github.com/aquiseb/lychee/types"
)

// Query contains the DB so it's made available in all resolvers
type Query struct {
	DB *db.Services
}

// NodeResolver contains the node for resolving
type NodeResolver struct {
	node types.Node
}

// UserResolver contains the DB and the model for resolving
type UserResolver struct {
	DB *db.Services
	m  models.User
}
