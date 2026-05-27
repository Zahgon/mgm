package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField struct contains a model's ID field.
type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// DateFields struct contains the `created_at` and `updated_at`
// fields that autofill when inserting or updating a model.
type DateFields struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Otherwise id must be ObjectId

// GetID method returns a model's ID
func (f *IDField) GetID() interface{} {
	_ = "STUB: not implemented"

	// SetID sets the value of a model's ID field.
	return nil
}

func (f *IDField) SetID(id interface{}) { _ = "STUB: not implemented"; return }

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook is used here to set the `created_at` field
// value when inserting a new model into the database.
// TODO: get context as param the next version(4).
func (f *DateFields) Creating() error { _ = "STUB: not implemented"; return nil }

// Saving hook is used here to set the `updated_at` field
// value when creating or updating a model.
// TODO: get context as param the next version(4).
func (f *DateFields) Saving() error { _ = "STUB: not implemented"; return nil }
