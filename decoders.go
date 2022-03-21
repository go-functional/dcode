package dcode

import (
	"encoding/json"
)

// Float64 returns a Decoder that can decode JSON directly
// into a float64
func Float64() Decoder[float64] {
	return DecoderFunc[float64](
		func(dc *json.Decoder) (float64, error) {
			return unmarshal[float64](dc)
		},
	)
}

// Bool returns a decoder that can decode JSON
// into a bool
func Bool() Decoder[bool] {
	return DecoderFunc[bool](
		func(dc *json.Decoder) (bool, error) {
			return unmarshal[bool](dc)
		},
	)
}

// Int returns a Decoder that can decode any JSON
// into an integer
func Int() Decoder[int] {
	return DecoderFunc[int](
		func(dc *json.Decoder) (int, error) {
			return unmarshal[int](dc)
		},
	)
}

type MapT map[string]interface{}

func Map() Decoder[MapT] {
	return DecoderFunc[MapT](
		func(dc *json.Decoder) (MapT, error) {
			return unmarshal[MapT](dc)
		},
	)
}

// String returns a Decoder that can decode JSON
// into a string
func String() Decoder[string] {
	return DecoderFunc[string](
		func(dc *json.Decoder) (string, error) {
			return unmarshal[string](dc)
		},
	)
}

type IntermediateT json.RawMessage

// Intermediate returns a Decoder that decodes JSON
// into a a yet-to-be-determined type.
func Intermediate() Decoder[IntermediateT] {
	return DecoderFunc[IntermediateT](
		func(dc *json.Decoder) (IntermediateT, error) {
			return unmarshal[IntermediateT](dc)
		},
	)
}
