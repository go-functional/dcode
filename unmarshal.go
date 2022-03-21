package dcode

import "encoding/json"

func unmarshal[T any](dc *json.Decoder) (T, error) {
	var ret T
	if err := dc.Decode(&ret); err != nil {
		return ret, err
	}
	return ret, nil
}
