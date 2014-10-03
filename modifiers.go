// Package divvy provides common collection operations for Go slices.
package divvy

import "reflect"

// Append ...
func (d Divvy) Append(items ...interface{}) Divvy {
	var values []reflect.Value

	for i := 0; i < len(items); i++ {
		var (
			value   = reflect.ValueOf(items[i])
			element = value.Type()
		)

		if element != d.element {
			d.err = InvalidElementError{Actual: element, Expected: d.element}
			return d
		}

		values = append(values, reflect.ValueOf(items[i]))
	}

	d.collection = reflect.Append(d.collection, values...)
	return d
}

// Prepend ...
func (d Divvy) Prepend(items ...interface{}) Divvy {
	var slice = reflect.MakeSlice(d.typ, 0, 0)

	for i := 0; i < len(items); i++ {
		var (
			value   = reflect.ValueOf(items[i])
			element = value.Type()
		)

		if element != d.element {
			d.err = InvalidElementError{Actual: element, Expected: d.element}
			return d
		}

		slice = reflect.Append(slice, value)
	}

	d.collection = reflect.AppendSlice(slice, d.collection)
	return d
}
