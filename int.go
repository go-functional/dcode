package dcode

import (
	"encoding/json"
)

// Int decodes any JSON field into an integer
func Int() Decoder {
	return newDecoder(func(b []byte) (interface{}, error) {
		var i int
		if err := json.Unmarshal(b, &i); err != nil {
			return 0, err
		}
		return i, nil
	})
}
