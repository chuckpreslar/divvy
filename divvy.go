// Package divvy provides common collection operations for Go slices.
package divvy

import "reflect"

// Divvy ...
type Divvy struct {
	collection reflect.Value
	typ        reflect.Type
	kind       reflect.Kind
	element    reflect.Type
	err        error
}

// Result ...
func (d Divvy) Result() (interface{}, error) {
	return d.collection.Interface(), d.err
}

// New ...
func New(v interface{}) Divvy {
	var (
		value = reflect.ValueOf(v)
		typ   = value.Type()
		kind  = value.Kind()
	)

	if reflect.Slice != kind {
		return Divvy{err: UnsupportedKindError{Kind: kind}}
	}

	var (
		element    = typ.Elem()
		collection = reflect.MakeSlice(typ, value.Len(), value.Len())
	)

	reflect.Copy(collection, value)

	return Divvy{collection: collection, typ: typ, kind: kind, element: element}
}
