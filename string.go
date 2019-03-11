package dcode

import (
	"encoding/json"
)

// String decodes any JSON field into a string
func String() Decoder {
	return newDecoder(func(b []byte) (interface{}, error) {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return "", err
		}
		return s, nil
	})
}
