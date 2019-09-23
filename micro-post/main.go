package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/resolvers"
	"github.com/astenmies/lychee/micro-post/schema"
	"github.com/spf13/viper"
)

func main() {
	port := viper.GetString("port")

	c, _ := core.GetClient()
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
