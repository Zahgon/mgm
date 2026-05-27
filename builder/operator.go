package builder

// Operator is an interface that should be implemented by structs used as operators.
type Operator interface {
	GetKey() string
	GetVal() interface{}
}

// BaseOperator is a simple operator struct that implementes the Operator interface.
type BaseOperator struct {
	key string
	val interface{}
}

// GetKey function returns the operator's key.
func (operator *BaseOperator) GetKey() string { _ = "STUB: not implemented"; return "" }

// GetVal function returns the operator's value.
func (operator *BaseOperator) GetVal() interface{} { _ = "STUB: not implemented"; return nil }

// New function creates a new base operator with the specified key and value
func New(key string, val interface{}) Operator { _ = "STUB: not implemented"; return *new(Operator) }

// Ensure that the BaseOperator implements the Operator interace
var _ Operator = &BaseOperator{}
