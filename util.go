package mgm

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Coll returns the collection associated with a model.
func Coll(m Model, opts ...*options.CollectionOptions) *Collection {
	_ = "STUB: not implemented"
	return nil
}

// CollName returns a model's collection name. The `CollectionNameGetter` will be used
// if the model implements this interface. Otherwise, the collection name is inferred
// based on the model's type using reflection.
func CollName(m Model) string { _ = "STUB: not implemented"; return "" }

// UpsertTrueOption returns new instance of UpdateOptions with the upsert property set to true.
func UpsertTrueOption() *options.UpdateOptions { _ = "STUB: not implemented"; return nil }
