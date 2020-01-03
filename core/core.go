package core

import (
	"io/ioutil"
	"net/http"

	"github.com/astenmies/lychee/core/static"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Graphql middleware renders the handler
func Graphql(schema string, resolvers interface{}) http.HandlerFunc {
	parsedSchema := graphql.MustParseSchema(schema, resolvers)

	// SHOULD PASS DB TO RESOLVERS
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relayHandler := relay.Handler{Schema: parsedSchema}
		relayHandler.ServeHTTP(w, r)
	})
}

// Playground middleware renders Graphiql
func Playground() http.HandlerFunc {
	graphiql, _ := static.Asset("static/index.html")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(graphiql)
	})
}

// GetSchema gets the schema at a provided path
func GetSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
