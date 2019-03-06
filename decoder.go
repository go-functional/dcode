package dcode

type Decoder struct {
	call func([]byte) (interface{}, error)
}

func newDecoder(fn func([]byte) (interface{}, error)) Decoder {
	return Decoder{call: fn}
}
