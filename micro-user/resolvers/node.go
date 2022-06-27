package resolvers

import (
	"context"

	"github.com/aquiseb/lychee/micro-user/models"
	"github.com/graph-gophers/graphql-go"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

// Node is the top level node query resolver for relationships between services
func (q *Query) Node(args struct{ ID string }) (*NodeResolver, error) {
	dbName := viper.GetString("db.name")

	// user := users[args.ID]
	var result models.User
	collection := q.DB.Client.Database(dbName).Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"id": args.ID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &NodeResolver{&UserResolver{
		m:  result,
		DB: q.DB, // Careful here, this is mandatory in order to query stuff on DB from within the UserResolver
	}}, nil
}

// ID of the node
func (n *NodeResolver) ID() graphql.ID {
	return n.node.ID()
}

// ToUser resolves the node to a user
func (n *NodeResolver) ToUser() (*UserResolver, bool) {
	user, ok := n.node.(*UserResolver)
	return user, ok
}
