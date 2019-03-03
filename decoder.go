package dcode

type JSONValue struct {
	data []byte
}

func NewJSONValue(b []byte) JSONValue {
	return JSONValue{data: b}
}

type Decoder func(JSONValue) (interface{}, error)

// func Field(name string, decoder Decoder) Decoder {

// }

func DecodeString(d Decoder, s string) (interface{}, error) {
	return DecodeBytes(d, []byte(s))
}

func DecodeBytes(d Decoder, b []byte) (interface{}, error) {
	jsonVal := NewJSONValue(b)
	return d(jsonVal)
}
