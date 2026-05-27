package mgm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config *Config
var client *mongo.Client
var db *mongo.Database

// Config struct contains extra configuration properties for the mgm package.
type Config struct {
	// Set to 10 second (10*time.Second) for example.
	CtxTimeout time.Duration
}

// NewCtx function creates and returns a new context with the specified timeout.
func NewCtx(timeout time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Ctx function creates and returns a new context with a default timeout value.
func Ctx() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func ctx() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

// NewClient returns a new mongodb client.
func NewClient(opts ...*options.ClientOptions) (*mongo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCollection returns a new collection with the supplied database.
func NewCollection(db *mongo.Database, name string, opts ...*options.CollectionOptions) *Collection {
	_ = "STUB: not implemented"
	return nil
}

// ResetDefaultConfig resets the configuration values, client and database.
func ResetDefaultConfig() { _ = "STUB: not implemented"; return }

// SetDefaultConfig initializes the client and database using the specified configuration values, or default.
func SetDefaultConfig(conf *Config, dbName string, opts ...*options.ClientOptions) (err error) {
	_ = "STUB: not implemented"

	// Use the predefined configuration values as default if the user
	// does not provide any.
	return nil
}

// CollectionByName returns a new collection using the current configuration values.
func CollectionByName(name string, opts ...*options.CollectionOptions) *Collection {
	_ = "STUB: not implemented"
	return nil
}

// DefaultConfigs returns the current configuration values, client and database.
func DefaultConfigs() (*Config, *mongo.Client, *mongo.Database, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// defaultConf are the default configuration values when none are provided
// to the `SetDefaultConfig` method.
func defaultConf() *Config { _ = "STUB: not implemented"; return nil }
