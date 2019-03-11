package dcode

// Decoder is the core type in dcode. This is basically a pure function
// that can decode ANY JSON value into a particular type.
//
// You can't implement one of these directly. Instead use one of the built-in
// ones like Int(), String(), etc... and build them up with Field(...) or
// First(...).Then(...).Into(...).
//
// Check out the documentation for Field() or Builder for more information
type Decoder struct {
	call func([]byte) (interface{}, error)
}

func newDecoder(fn func([]byte) (interface{}, error)) Decoder {
	return Decoder{call: fn}
}
