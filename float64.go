package dcode

import (
	"encoding/json"
)

// Float64 decodes any JSON field into a float64
func Float64() Decoder {
	return newDecoder(func(b []byte) (interface{}, error) {
		var f float64
		if err := json.Unmarshal(b, &f); err != nil {
			return "", err
		}
		return f, nil
	})
}
