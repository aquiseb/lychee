module github.com/astenmies/lychee

go 1.12

require (
	github.com/carted/graphql v0.7.6
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.7.0 // indirect
	github.com/go-siris/siris v7.4.0+incompatible // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/graph-gophers/graphql-go v0.0.0-20190917030536-38a077bc812d
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/nautilus/gateway v0.0.14
	github.com/nautilus/graphql v0.0.7
	github.com/oxequa/interact v0.0.0-20171114182912-f8fb5795b5d7 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli/v2 v2.0.0 // indirect
	github.com/valyala/fasttemplate v1.1.0 // indirect
	github.com/vektah/gqlparser v1.1.2
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	github.com/yyh-gl/realize v2.1.6+incompatible // indirect
	go.mongodb.org/mongo-driver v1.1.1
	golang.org/x/text v0.3.2 // indirect
)

// replace github.com/carted/graphql => github.com/astenmies/graph-go v0.7.7 // indirect

replace (
	github.com/astenmies/lychee => ./
	github.com/oxequa/realize => github.com/yyh-gl/realize v2.1.6+incompatible // indirect
	github.com/urfave/cli.V2 => github.com/urfave/cli/v2 v2.0.0 // indirect
)
