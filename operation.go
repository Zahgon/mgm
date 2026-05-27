package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(ctx context.Context, c *Collection, model Model, opts ...*options.InsertOneOptions) error {
	_ = "STUB: not implemented"
	// Call to saving hook
	return nil
}

// Set new id

func first(ctx context.Context, c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func update(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	_ = "STUB: not implemented"
	// Call to saving hook
	return nil
}

func del(ctx context.Context, c *Collection, model Model) error {
	_ = "STUB: not implemented"
	return nil
}
