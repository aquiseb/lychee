module github.com/astenmies/lychee

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/graph-gophers/graphql-go v0.0.0-20190917030536-38a077bc812d
	github.com/nautilus/gateway v0.0.14
	github.com/nautilus/graphql v0.0.7
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.4.0
	github.com/vektah/gqlparser v1.1.2
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.1
)

// replace github.com/carted/graphql => github.com/astenmies/graph-go v0.7.7 // indirect

replace github.com/astenmies/lychee => ./
