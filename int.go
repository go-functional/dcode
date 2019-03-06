package dcode

import (
	"encoding/json"
)

func Int() Decoder {
	return func(b []byte) (interface{}, error) {
		var i int
		if err := json.Unmarshal(b, &i); err != nil {
			return 0, err
		}
		return i, nil
	}
}
