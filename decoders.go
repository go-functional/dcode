package dcode

import (
	"encoding/json"
)

// Float64 returns a Decoder that can decode JSON directly
// into a float64
func Float64() Decoder[float64] {
	var ret float64
	return DecoderFunc[float64](
		func(b []byte) (float64, error) {
			if err := json.Unmarshal(b, &ret); err != nil {
				return ret, nil
			}
			return ret, nil
		},
	)
}

// Bool returns a decoder that can decode JSON
// into a bool
func Bool() Decoder[bool] {
	return DecoderFunc[bool](func(b []byte) (bool, error) {
		var ret bool
		if err := json.Unmarshal(b, &ret); err != nil {
			return false, err
		}
		return ret, nil
	})
}

// Int returns a Decoder that can decode any JSON
// into an integer
func Int() Decoder[int] {
	var zero int
	return DecoderFunc[int](
		func(b []byte) (int, error) {
			var ret int
			if err := json.Unmarshal(b, &ret); err != nil {
				return zero, err
			}
			return ret, nil
		},
	)
}

type MapT map[string]interface{}

func Map() Decoder[MapT] {
	zero := map[string]interface{}{}
	return DecoderFunc[MapT](
		func(b []byte) (MapT, error) {
			var ret map[string]interface{}
			if err := json.Unmarshal(b, &ret); err != nil {
				return zero, nil
			}
			return ret, nil
		},
	)
}

// String returns a Decoder that can decode JSON
// into a string
func String() Decoder[string] {
	var zero string
	return DecoderFunc[string](
		func(b []byte) (string, error) {
			var ret string
			if err := json.Unmarshal(b, &ret); err != nil {
				return zero, err
			}
			return ret, nil
		},
	)
}
