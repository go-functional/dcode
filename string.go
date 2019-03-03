package dcode

import (
	"encoding/json"
)

func String() Decoder {
	return func(val JSONValue) (interface{}, error) {
		var ret string
		if err := json.Unmarshal(val.data, &ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
