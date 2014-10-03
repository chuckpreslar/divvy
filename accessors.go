// Package divvy provides common collection operations for Go slices.
package divvy

import "reflect"

// AtIndex ...
func (d Divvy) AtIndex(index int) (interface{}, error) {
	if nil != d.err {
		return reflect.Zero(d.element).Interface(), d.err
	}

	var length, err = d.Length()

	if nil != err {
		return reflect.Zero(d.element).Interface(), d.err
	}

	if 0 > index || length-1 < index {
		d.err = InvalidIndexError{Index: index, Length: length}
		return reflect.Zero(d.element).Interface(), d.err
	}

	return d.collection.Index(index).Interface(), d.err
}
