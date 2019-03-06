package dcode

import (
	"encoding/json"
	"fmt"
)

func Field(name string, decoder Decoder) Decoder {
	return func(b []byte, i interface{}) error {
		// if we're at a leaf, things should just decode!
		leafErr := json.Unmarshal(b, i)
		if leafErr == nil {
			return nil
		}
		// otherwise we're not at a leaf, so traverse down the tree.
		//
		// decode what we've got into a map, then get the object under
		// name, and call the decoder on that
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return err
		}
		next, ok := m[name]
		if !ok {
			return fmt.Errorf("key %s not found", name)
		}
		return decoder(b, next)
	}
}
