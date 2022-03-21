package dcode

import (
	"bytes"
	"encoding/json"
)

// Decoder is the core type in dcode. This is basically a pure function
// that can decode ANY JSON value into a particular type.
//
// You can't implement one of these directly. Instead use one of the built-in
// ones like Int(), String(), etc... and build them up with Field(...) or
// First(...).Then(...).Into(...).
//
// Check out the documentation for Field() or Builder for more information
type Decoder[T any] interface {
	Decode(*json.Decoder) (T, error)
}

type DecoderFunc[T any] func(dc *json.Decoder) (T, error)

func (d DecoderFunc[T]) Decode(dc *json.Decoder) (T, error) {
	return d(dc)
}

// Decode decodes b into a type T using d
func DecodeBytes[T any](b []byte, d Decoder[T]) (T, error) {
	return d.Decode(
		json.NewDecoder(bytes.NewBuffer(b)),
	)
}
