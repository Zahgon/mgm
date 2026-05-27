package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection performs operations on models and the given Mongodb collection
type Collection struct {
	*mongo.Collection
}

// FindByID method finds a doc and decodes it to a model, otherwise returns an error.
// The id field can be any value that if passed to the `PrepareID` method, it returns
// a valid ID (e.g string, bson.ObjectId).
func (coll *Collection) FindByID(id interface{}, model Model, opts ...*options.FindOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// FindByIDWithCtx method finds a doc and decodes it to a model, otherwise returns an error.
// The id field can be any value that if passed to the `PrepareID` method, it returns
// a valid ID (e.g string, bson.ObjectId).
func (coll *Collection) FindByIDWithCtx(ctx context.Context, id interface{}, model Model, opts ...*options.FindOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// First method searches and returns the first document in the search results.
func (coll *Collection) First(filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// FirstWithCtx method searches and returns the first document in the search results.
func (coll *Collection) FirstWithCtx(ctx context.Context, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Create method inserts a new model into the database.
func (coll *Collection) Create(model Model, opts ...*options.InsertOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateWithCtx method inserts a new model into the database.
func (coll *Collection) CreateWithCtx(ctx context.Context, model Model, opts ...*options.InsertOneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Update function persists the changes made to a model to the database.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (coll *Collection) Update(model Model, opts ...*options.UpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateWithCtx function persists the changes made to a model to the database using the specified context.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (coll *Collection) UpdateWithCtx(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete method deletes a model (doc) from a collection.
// To perform additional operations when deleting a model
// you should use hooks rather than overriding this method.
func (coll *Collection) Delete(model Model) error { _ = "STUB: not implemented"; return nil }

// DeleteWithCtx method deletes a model (doc) from a collection using the specified context.
// To perform additional operations when deleting a model
// you should use hooks rather than overriding this method.
func (coll *Collection) DeleteWithCtx(ctx context.Context, model Model) error {
	_ = "STUB: not implemented"
	return nil
}

// SimpleFind finds, decodes and returns the results.
func (coll *Collection) SimpleFind(results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// SimpleFindWithCtx finds, decodes and returns the results using the specified context.
func (coll *Collection) SimpleFindWithCtx(ctx context.Context, results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------------
// Aggregation methods
// --------------------------------

// SimpleAggregateFirst is just same as SimpleAggregateFirstWithCtx, but doesn't get context param.
func (coll *Collection) SimpleAggregateFirst(result interface{}, stages []interface{}, opts ...*options.AggregateOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SimpleAggregateFirstWithCtx performs a simple aggregation, decodes the first aggregate result and returns it using the provided result parameter.
// The value of `stages` can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (coll *Collection) SimpleAggregateFirstWithCtx(ctx context.Context, result interface{}, stages []interface{}, opts ...*options.AggregateOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SimpleAggregate is just same as SimpleAggregateWithCtx, but doesn't get context param.
func (coll *Collection) SimpleAggregate(results interface{}, stages []interface{}, opts ...*options.AggregateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// SimpleAggregateWithCtx performs a simple aggregation, decodes the aggregate result and returns the list using the provided result parameter.
// The value of `stages` can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (coll *Collection) SimpleAggregateWithCtx(ctx context.Context, results interface{}, stages []interface{}, opts ...*options.AggregateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// SimpleAggregateCursor is just same as SimpleAggregateCursorWithCtx, but
// doesn't get context.
func (coll *Collection) SimpleAggregateCursor(stages []interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SimpleAggregateCursorWithCtx performs a simple aggregation and returns a cursor over the resulting documents.
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (coll *Collection) SimpleAggregateCursorWithCtx(ctx context.Context, stages []interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
