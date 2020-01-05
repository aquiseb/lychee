package resolvers

import (
	"github.com/astenmies/lychee/micro-user/db"
	"github.com/astenmies/lychee/micro-user/models"
	"github.com/astenmies/lychee/types"
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
