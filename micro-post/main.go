package main

import (
	"fmt"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/resolvers"
	"github.com/astenmies/lychee/micro-post/schema"
)

func main() {
	c, _ := core.GetClient()
	schem := schema.GetSchema()
	database := &db.Services{c}

	r := &resolvers.Query{
		DB: database,
	}

	http.Handle("/graphql", core.Graphql(schem, r))
	http.Handle("/", core.Playground())

	// https://stackoverflow.com/a/48250354/9077800
	done := make(chan bool)
	go http.ListenAndServe(":4002", nil)
	fmt.Println("FEDERATION_SIGNAL_OK", "Started server on http://localhost:4002")
	<-done
}

func init() {
	// core.InitViper()
}
