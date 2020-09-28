package dcode

import (
	tree "github.com/bmatsuo/go-jsontree"
)

// Exists checks whether the value that decoder points to (in the JSON bytes)
// exists and is the expected type
func Exists(decoder Decoder, bytes []byte) bool {
	t := tree.New()
	if err := t.UnmarshalJSON(bytes); err != nil {
		return false
	}
	if _, err := decoder.call(t); err != nil {
		return false
	}
	return true
}
