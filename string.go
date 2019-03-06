package dcode

import (
	"encoding/json"
)

func String() Decoder {
	return newDecoder(func(b []byte) (interface{}, error) {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return "", err
		}
		return s, nil
	})
}
