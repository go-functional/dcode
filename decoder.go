package dcode

// Decoder is the core type in dcode. This is basically a pure function
// that can decode ANY JSON value into a particular type.
//
// You can't implement one of these directly. Instead use one of the built-in
// ones like Int(), String(), etc... and build them up with Field(...) or
// First(...).Then(...).Into(...).
//
// Check out the documentation for Field() or Builder for more information
type Decoder[T any] interface {
	Decode([]byte) (T, error)
}

type DecoderFunc[T any] func([]byte) (T, error)

func (d DecoderFunc[T]) Decode(b []byte) (T, error) {
	return d(b)
}
