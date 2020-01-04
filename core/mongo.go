package core

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetMongoClient returns a MongoDB Client
func GetMongoClient() (*mongo.Client, error) {
	uri := viper.GetString("db.uri")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	fmt.Println("Mongo connected")

	return client, nil
}

// FindOrSeed checks if the database was already seeded or seeds it
func FindOrSeed(collection *mongo.Collection, key string, data []interface{}) {
	dbName := viper.GetString("db.name")

	// Get a random element in data
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(data)
	filter := data[n]

	// Find the random element in the database
	var result interface{}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// Insert the data only if the random element was not found in the database
		if strings.Contains(err.Error(), "no documents in result") {
			_, err = collection.InsertMany(context.TODO(), data)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("DB", dbName, "contains", len(data), "demo", key)

}

// Seed inserts some demo data into the database
func Seed(c *mongo.Client) {
	dbName := viper.GetString("db.name")
	col := c.Database(dbName)

	dataToSeed := map[string][]interface{}{
		"users": {
			bson.M{"id": "1", "firstname": "Aline", "lastname": "MQ", "canWriteReviews": true},
			bson.M{"id": "2", "firstname": "Rose", "lastname": "DS", "canWriteReviews": true},
			bson.M{"id": "3", "firstname": "Seb", "lastname": "DS", "canWriteReviews": false},
		},
		"posts": {
			bson.M{"id": "1", "title": "First post", "slug": "first-post"},
			bson.M{"id": "2", "title": "Second post", "slug": "second-post"},
			bson.M{"id": "3", "title": "Third post", "slug": "third-post"},
		},
		"reviews": {
			bson.M{"id": "1", "stars": 5, "postId": "1", "userId": "3"},
			bson.M{"id": "2", "stars": 4, "postId": "2", "userId": "1"},
			bson.M{"id": "3", "stars": 3, "postId": "1", "userId": "2"},
			bson.M{"id": "4", "stars": 1, "postId": "1", "userId": "1"},
		},
	}

	for key, data := range dataToSeed {
		collection := col.Collection(key)
		FindOrSeed(collection, key, data)
	}

}
