package dcode

import (
	"encoding/json"
	"fmt"
)

func Field(name string, decoder Decoder) Decoder {
	return func(b []byte) (interface{}, error) {
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		iface, ok := m[name]
		if !ok {
			// TODO: use a standardized error
			return nil, fmt.Errorf("Field %s not found", name)
		}
		recoded, err := json.Marshal(iface)
		if err != nil {
			return nil, err
		}
		return decoder(recoded)
	}
}
