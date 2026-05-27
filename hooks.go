package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreatingHook is called before saving a new model to the database
// Deprecated: please use CreatingHookWithCtx
type CreatingHook interface {
	Creating() error
}

// CreatingHookWithCtx is called before saving a new model to the database
type CreatingHookWithCtx interface {
	Creating(context.Context) error
}

// CreatedHook is called after a model has been created
// Deprecated: Please use CreatedHookWithCtx
type CreatedHook interface {
	Created() error
}

// CreatedHookWithCtx is called after a model has been created
type CreatedHookWithCtx interface {
	Created(context.Context) error
}

// UpdatingHook is called before updating a model
// Deprecated: Please use UpdatingHookWithCtx
type UpdatingHook interface {
	Updating() error
}

// UpdatingHookWithCtx is called before updating a model
type UpdatingHookWithCtx interface {
	Updating(context.Context) error
}

// UpdatedHook is called after a model is updated
// Deprecated: Please use UpdatedHookWithCtx
type UpdatedHook interface {
	// Deprecated:
	Updated(result *mongo.UpdateResult) error
}

// UpdatedHookWithCtx is called after a model is updated
type UpdatedHookWithCtx interface {
	Updated(ctx context.Context, result *mongo.UpdateResult) error
}

// SavingHook is called before a model (new or existing) is saved to the database.
// Deprecated: Please use SavingHookWithCtx
type SavingHook interface {
	Saving() error
}

// SavingHookWithCtx is called before a model (new or existing) is saved to the database.
type SavingHookWithCtx interface {
	Saving(context.Context) error
}

// SavedHook is called after a model is saved to the database.
// Deprecated: Please use SavedHookWithCtx
type SavedHook interface {
	Saved() error
}

// SavedHookWithCtx is called after a model is saved to the database.
type SavedHookWithCtx interface {
	Saved(context.Context) error
}

// DeletingHook is called before a model is deleted
// Deprecated: Please use DeletingHookWithCtx
type DeletingHook interface {
	Deleting() error
}

// DeletingHookWithCtx is called before a model is deleted
type DeletingHookWithCtx interface {
	Deleting(context.Context) error
}

// DeletedHook is called after a model is deleted
// Deprecated: Please use DeletedHookWithCtx
type DeletedHook interface {
	Deleted(result *mongo.DeleteResult) error
}

// DeletedHookWithCtx is called after a model is deleted
type DeletedHookWithCtx interface {
	Deleted(ctx context.Context, result *mongo.DeleteResult) error
}

func callToBeforeCreateHooks(ctx context.Context, model Model) error {
	_ = "STUB: not implemented"
	return nil
}

func callToBeforeUpdateHooks(ctx context.Context, model Model) error {
	_ = "STUB: not implemented"
	return nil
}

func callToAfterCreateHooks(ctx context.Context, model Model) error {
	_ = "STUB: not implemented"
	return nil
}

func callToAfterUpdateHooks(ctx context.Context, updateResult *mongo.UpdateResult, model Model) error {
	_ = "STUB: not implemented"
	return nil
}

func callToBeforeDeleteHooks(ctx context.Context, model Model) error {
	_ = "STUB: not implemented"
	return nil
}

func callToAfterDeleteHooks(ctx context.Context, deleteResult *mongo.DeleteResult, model Model) error {
	_ = "STUB: not implemented"
	return nil
}
