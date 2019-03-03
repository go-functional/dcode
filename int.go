package dcode

import (
	"encoding/json"
)

func Int() Decoder {
	return func(val JSONValue) (interface{}, error) {
		var ret int
		if err := json.Unmarshal(val.data, &ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
