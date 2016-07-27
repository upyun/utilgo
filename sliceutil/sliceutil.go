// Package sliceutil contains some slice related helper functions.
package sliceutil

// GetSliceIndex returns a's index from list
func GetSliceIndex(a interface{}, list []interface{}) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}
