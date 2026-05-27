package util

import (
	"regexp"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase returns snake_case of the provided value.
func ToSnakeCase(str string) string { _ = "STUB: not implemented"; return "" }
