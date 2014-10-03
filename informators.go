// Package divvy provides common collection operations for Go slices.
package divvy

// Length ...
func (d Divvy) Length() (int, error) {
	return d.collection.Len(), d.err
}

// Count ...
func (d Divvy) Count() (int, error) {
	return d.Length()
}

// Size ...
func (d Divvy) Size() (int, error) {
	return d.Length()
}

// IsEmpty ...
func (d Divvy) IsEmpty() (bool, error) {
	var length, _ = d.Length()
	return (length == 0), d.err
}
