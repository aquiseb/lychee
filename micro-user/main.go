package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-user/db"
	resolvers "github.com/astenmies/lychee/micro-user/resolvers"
	"github.com/astenmies/lychee/micro-user/schema"
)

// https://github.com/graph-gophers/graphql-go/issues/106#issuecomment-350231819
// RootResolver is extended with each "microservice" resolver

func main() {
	port := viper.GetString("port")

	c, _ := core.GetMongoClient()
	core.Seed(c) // seed demo data

	schem := schema.GetSchema()
	database := &db.Services{c}

	r := &resolvers.Query{
		DB: database,
	}

	http.Handle("/graphql", core.Graphql(schem, r))
	http.Handle("/", core.Playground())

	fmt.Println("Starting server on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	core.InitViper()
}
