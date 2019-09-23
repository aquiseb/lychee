package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/astenmies/lychee/core"
	resolvers "github.com/astenmies/lychee/micro-hello/resolvers"
	"github.com/astenmies/lychee/micro-hello/schema"
)

// https://github.com/graph-gophers/graphql-go/issues/106#issuecomment-350231819
// RootResolver is extended with each "microservice" resolver

// GetSchema returns the schema of Post
func GetSchema() string {
	// OPTION 1 - string schema
	// s := `
	// 	schema {
	// 			query: Query
	// 	}
	// 	type Query {
	// 			getGreeting: String!
	// 	}
	// `

	// OPTION 2 - Read from file
	// s, _ := core.GetSchema("schema/schema.graphql")

	// OPTION 3 - Use go-bindata
	s, _ := schema.Asset("schema/schema.graphql")
	stringSchema := string(s)

	return stringSchema
}

func main() {
	port := viper.GetString("port")
	s := GetSchema()
	r := &resolvers.Query{}

	http.Handle("/graphql", core.Graphql(s, r))
	http.Handle("/", core.Playground())

	fmt.Println("Starting server on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	core.InitViper()
}
