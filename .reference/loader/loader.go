package loader

import (
	"context"
	"fmt"

	"github.com/nicksrandall/dataloader"
)

// The key type is created so that these values do not collide with other keys that some other
// package may have placed on the context.
//
// For more explanation of this pattern, see:
// https://medium.com/@matryer/context-keys-in-go-5312346a868d
type key string

const (
	articleLoaderKey key = "article"
)

type Client interface {
	articleGetter
}

// Initialize a lookup map of context keys to batch functions.
//
// When Attach is called on the Collection, the batch functions are used to create new dataloader
// instances. The instances are attached to the request context at the provided keys.
//
// The keys are then used to extract the dataloader instances from the request context.
func Initialize(client Client) Collection {
	return Collection{
		lookup: map[key]dataloader.BatchFunc{
			articleLoaderKey: newArticleLoader(client),
		},
	}
}

// Collection holds an internal lookup of initialized batch data load functions.
type Collection struct {
	lookup map[key]dataloader.BatchFunc
}

// Attach creates new instances of dataloader.Loader and attaches the instances on the request context.
// We do this because the dataloader has an in-memory cache that is scoped to the request.
func (c Collection) Attach(ctx context.Context) context.Context {
	for k, batchFn := range c.lookup {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFn))
	}

	return ctx
}

// extract is a helper function to make common get-value, assert-type, return-error-or-value
// operations easier.
func extract(ctx context.Context, k key) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", k)
	}

	return ldr, nil
}

// Implements Stringer.
func (k key) String() string {
	return string(k)
}
