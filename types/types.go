package types

import (
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	*mongo.Client
}

// Node interface is required by nautilus gateway
// See here The Node interface here: https://itnext.io/a-guide-to-graphql-schema-federation-part-1-995b639ac035?
// The only requirement that schema federation imposes on a service is that it must satisfy the Relay Global Object Identification Specification.
// See here: https://facebook.github.io/relay/graphql/objectidentification.htm
type Node interface {
	ID() graphql.ID
}
