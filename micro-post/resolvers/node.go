package resolvers

import (
	"context"

	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

// var users = map[string]*models.User{
// 	"1": {
// 		ID:       "1",
// 		LastName: "Mies",
// 	},
// }

func (n *NodeResolver) ID() graphql.ID {
	return n.node.ID()
}

func (n *NodeResolver) ToUser() (*UserResolver, bool) {
	user, ok := n.node.(*UserResolver)
	return user, ok
}

func (n *NodeResolver) ToPost() (*PostResolver, bool) {
	post, ok := n.node.(*PostResolver)
	return post, ok
}

func (n *NodeResolver) ToReview() (*ReviewResolver, bool) {
	review, ok := n.node.(*ReviewResolver)
	return review, ok
}

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
