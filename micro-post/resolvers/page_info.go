package resolvers

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// StartCursor is the first id of the results, useful for pagination
func (p *PageInfo) StartCursor(ctx context.Context) *graphql.ID {
	return &p.startCursor
}

// EndCursor is the first id of the results, useful for pagination
func (p *PageInfo) EndCursor(ctx context.Context) *graphql.ID {
	return &p.endCursor
}

// HasNextPage tells if there's a next page when doing pagination
func (p *PageInfo) HasNextPage(ctx context.Context) bool {
	return p.hasNextPage
}

// HasPreviousPage tells if there's a previous page when doing pagination
func (p *PageInfo) HasPreviousPage(ctx context.Context) bool {
	return p.hasPreviousPage
}
