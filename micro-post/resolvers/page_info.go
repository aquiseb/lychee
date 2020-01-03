package resolvers

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// StartCursor ...
func (p *PageInfo) StartCursor(ctx context.Context) *graphql.ID {
	return &p.startCursor
}

// EndCursor ...
func (p *PageInfo) EndCursor(ctx context.Context) *graphql.ID {
	return &p.endCursor
}

// HasNextPage ...
func (p *PageInfo) HasNextPage(ctx context.Context) bool {
	return p.hasNextPage
}

// HasPreviousPage ...
func (p *PageInfo) HasPreviousPage(ctx context.Context) bool {
	return p.hasPreviousPage
}
