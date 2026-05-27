package mgm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// TransactionFunc is a handler to manage a transaction.
type TransactionFunc func(session mongo.Session, sc mongo.SessionContext) error

// Transaction creates a transaction with the default client.
func Transaction(f TransactionFunc) error { _ = "STUB: not implemented"; return nil }

// TransactionWithCtx creates a transaction with the given context and the default client.
func TransactionWithCtx(ctx context.Context, f TransactionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// TransactionWithClient creates a transaction with the given client.
func TransactionWithClient(ctx context.Context, client *mongo.Client, f TransactionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

//start session need to get options.

// startTransaction need to get options.
