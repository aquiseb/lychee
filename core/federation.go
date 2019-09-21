package core

import (
	"fmt"
	"net/http"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
)

var Version string

func Federation() {
	fmt.Printf(`Version: %s`, Version)

	schemas, err := graphql.IntrospectRemoteSchemas(
		"http://localhost:4001/graphql",
		"http://localhost:4002/graphql",
	)
	if err != nil {
		panic(err)
	}

	// create the gateway instance
	gw, err := gateway.New(schemas)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", gw.GraphQLHandler)
	http.HandleFunc("/", gw.PlaygroundHandler)

	fmt.Println("Starting server on http://localhost:4000")
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
