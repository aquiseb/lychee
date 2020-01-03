package db

import (
	"context"
	"fmt"
	"time"

	"github.com/astenmies/lychee/micro-post/models"
	"github.com/astenmies/lychee/types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// Services expose db methods
type Services types.DB

// Check just prints a string
func (s *Services) Check(str string) {
	fmt.Printf("db methods are working! %s", str)
}

// GetPostById retrieves a post from db based on id
func (s *Services) GetPostById(filter bson.M) (*models.Post, error) {
	var result models.Post
	collection := s.Client.Database("lychee").Collection("posts")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, errors.Cause(err)
	}

	return &result, nil
}

// GetReviewById retrieves a review from db based on id
func (s *Services) GetReviewById(filter bson.M) (*models.Review, error) {
	var result models.Review
	collection := s.Client.Database("lychee").Collection("reviews")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, errors.Cause(err)
	}

	return &result, nil
}

func (s *Services) GetReviewsByPostId(filter bson.M) (*[]*models.Review, error) {
	var results []*models.Review

	collection := s.Client.Database("lychee").Collection("reviews")

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, errors.Cause(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var result models.Review
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

func (s *Services) GetReviewsByUserId(filter bson.M) (*[]*models.Review, error) {
	var results []*models.Review

	collection := s.Client.Database("lychee").Collection("reviews")

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, errors.Cause(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var result models.Review
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
