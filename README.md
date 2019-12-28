[![Lychee](https://repository-images.githubusercontent.com/210030187/600c7380-dcb9-11e9-992b-bbe7a38e48e5)](https://github.com/astenmies/lychee)

<h1 align="center">
lychee
</h1>

<p align="center">
  <strong>
    A simple example of GraphQL microservices federated in Golang, consumed by a Next.js client.
  </strong>
</p>

---

## This repository is a work in progress...

## Setup

```shell
go get -u github.com/astenmies/lychee
go get ./...
```

## Start

Open each microservice in a separate terminal tab. Start each of them with

```shell
cd micro-xxxxx
./start.sh
```

The federation of each of these microservices is done with the `micro-federation` module.

```shell
cd micro-federation
go run main.go
```

## Failed downloading ...

https://github.com/oxequa/realize/issues/253#issuecomment-532045314

or

```
go get -v gopkg.in/urfave/cli.v2 \
&& go get -v github.com/oxequa/realize
```

## Configuration

The configuration is handled by [viper](github.com/spf13/viper).
Add a `_config` file in your microservice.

## Versioning

Versioning is handled by [govvv](github.com/ahmetb/govvv).
To change the version, edit the `VERSION` file.

## Build

Each microservice must be built separately.
Add a `VERSION` file and specify the version of the microservice.

```shell
govvv build
```

## Static assets

**lychee** uses [go-bindata](https://github.com/jteeuwen/go-bindata) to convert any file into managable Go source code.
Useful for embedding binary data (like GraphQL schemas) into a go program.

Write your microservice schema package, then generate `bindata.go`.

```shell
go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
```

Here's how to use it from within main.go.

```go
// GetSchema returns the schema of Post
func GetSchema() string {
	s, _ := schema.Asset("schema/schema.graphql")
	stringSchema := string(s)

	return stringSchema
}
```

## Graphql Queries

```graphql
{
    post(id: "2") {
        title
        reviews {
        edges {
            node {
                id
                stars
            }
        }
    }
}
```

```graphql
{
  allUsers {
    id
    firstName
  }
}
```
