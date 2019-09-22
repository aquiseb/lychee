module github.com/astenmies/lychee

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/graph-gophers/graphql-go v0.0.0-20190917030536-38a077bc812d
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/nautilus/gateway v0.0.14
	github.com/nautilus/graphql v0.0.7
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.4.0
	github.com/vektah/gqlparser v1.1.2
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.1
	golang.org/x/sys v0.0.0-20190921190940-14da1ac737cc // indirect
	golang.org/x/text v0.3.2 // indirect
)

// replace github.com/carted/graphql => github.com/astenmies/graph-go v0.7.7 // indirect

replace github.com/astenmies/lychee => ./
