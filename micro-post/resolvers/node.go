package resolvers

import (
	"context"

	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

type Node interface {
	ID() graphql.ID
}

type NodeResolver struct {
	node Node
}

var users = map[string]*models.User{
	"1": {
		ID:       "1",
		LastName: "Mies",
	},
}

func (n *NodeResolver) ID() graphql.ID {
	return n.node.ID()
}

func (n *NodeResolver) ToUser() (*UserResolver, bool) {
	user, ok := n.node.(*UserResolver)
	return user, ok
}

func (r *Query) Node(args struct{ ID string }) (*NodeResolver, error) {
	dbName := viper.GetString("db.name")

	user := users[args.ID]
	var result models.Post
	collection := r.DB.Client.Database(dbName).Collection("user")
	err := collection.FindOne(context.TODO(), bson.M{"id": args.ID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &NodeResolver{&UserResolver{*user}}, nil
}
