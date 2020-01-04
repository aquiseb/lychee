package db

import (
	"context"

	"github.com/astenmies/lychee/micro-post/models"
	"github.com/astenmies/lychee/types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// Services expose db methods
type Services types.DB

// GetUserByID retrieves a user from db based on id
func (s *Services) GetUserByID(filter bson.M) (*models.Post, error) {
	var result models.Post
	collection := s.Client.Database("lychee").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, errors.Cause(err)
	}

	return &result, nil
}
