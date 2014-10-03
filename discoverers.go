// Package divvy provides common collection operations for Go slices.
package divvy

import "reflect"

// IndexOf ...
func (d Divvy) IndexOf(item interface{}) (int, error) {
	if nil != d.err {
		return -1, d.err
	}

	for i := 0; i < d.collection.Len(); i++ {
		var value = d.collection.Index(i).Interface()
		if reflect.DeepEqual(value, item) {
			return i, d.err
		}
	}

	return -1, d.err
}

// LastIndexOf ...
func (d Divvy) LastIndexOf(item interface{}) (int, error) {
	if nil != d.err {
		return -1, d.err
	}

	for i := d.collection.Len() - 1; i >= 0; i-- {
		var value = d.collection.Index(i).Interface()
		if reflect.DeepEqual(value, item) {
			return i, d.err
		}
	}

	return -1, d.err
}

// Contains ...
func (d Divvy) Contains(item interface{}) (bool, error) {
	var index, err = d.IndexOf(item)

	if nil != err {
		return false, err
	} else if -1 == index {
		return false, err
	}

	return true, err
}
