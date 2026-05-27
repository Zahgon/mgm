package mgm

import (
	"strings"
)

var escape = strings.NewReplacer("$", "\uFF04", ".", "\uFF0E")
var unescape = strings.NewReplacer("\uFF04", "$", "\uFF0E", ".")

// Escape escapes the mongo key for . and $ characters.
func Escape(key string) string { _ = "STUB: not implemented"; return "" }

// Unescape unescapes the mongo key for . and $ characters.
func Unescape(key string) string { _ = "STUB: not implemented"; return "" }
