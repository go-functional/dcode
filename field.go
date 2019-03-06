package dcode

import (
	"encoding/json"
	"fmt"
)

func Field(name string, decoder Decoder) Decoder {
	return newDecoder(func(b []byte) (interface{}, error) {
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		iface, ok := m[name]
		if !ok {
			// TODO: use a standardized error
			return nil, fmt.Errorf("Field %s not found", name)
		}
		// we need to re-encode the sub-object so we can pass those bytes
		// down to the next decoder.
		//
		// TODO: figure out a better way to do this
		recoded, err := json.Marshal(iface)
		if err != nil {
			return nil, err
		}
		return decoder.call(recoded)
	})
}
