package dcode

import (
	"encoding/json"
)

// Bool decodes any JSON field into a bool
func Bool() Decoder {
	return newDecoder(func(b []byte) (interface{}, error) {
		var bl bool
		if err := json.Unmarshal(b, &bl); err != nil {
			return "", err
		}
		return bl, nil
	})
}
