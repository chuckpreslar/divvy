// Package divvy provides common collection operations for Go slices.
package divvy

import (
	"fmt"
	"reflect"
)

// InvalidElementError ...
type InvalidElementError struct {
	Actual   reflect.Type
	Expected reflect.Type
}

// InvalidElementError ...
func (i InvalidElementError) Error() string {
	return fmt.Sprintf("cannot use value with type `%s` when expected was `%s`", i.Actual, i.Expected)
}

// InvalidIndexError ...
type InvalidIndexError struct {
	Index  int
	Length int
}

// Error ...
func (i InvalidIndexError) Error() string {
	return fmt.Sprintf("attempt to access invalid index of `%d` within collection of length `%d`", i.Index, i.Length)
}

// UnsupportedKindError occurs when a Serializer encounters
// a reflect.Kind it does not have the ability to interact
// with.
type UnsupportedKindError struct {
	Kind reflect.Kind
}

// Error implements the `error` interface for the
// UnsupportedKindError type.
func (u UnsupportedKindError) Error() string {
	return fmt.Sprintf("encountered a value with reflect.Kind of `%s`, only `%s` is supported", u.Kind, reflect.Slice)
}
