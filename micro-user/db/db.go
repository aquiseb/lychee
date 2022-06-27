package db

import (
	"context"
	"time"

	"github.com/aquiseb/lychee/micro-user/models"
	"github.com/aquiseb/lychee/types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// Services expose db methods
type Services types.DB

// GetUserByID retrieves a user from db based on id
func (s *Services) GetUserByID(filter bson.M) (*models.User, error) {
	var result models.User
	collection := s.Client.Database("lychee").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, errors.Cause(err)
	}

	return &result, nil
}

func (s *Services) GetAllUsers() (*[]*models.User, error) {
	var results []*models.User

	collection := s.Client.Database("lychee").Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, errors.Cause(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var result models.User
		err := cursor.Decode(&result)
		if err != nil {
			return nil, errors.Cause(err)
		}
		// do something with the result
		results = append(results, &result)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Cause(err)
	}

	//dont forget to close the cursor
	defer cursor.Close(context.TODO())

	return &results, nil
}
