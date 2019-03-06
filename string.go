package dcode

import (
	"encoding/json"
)

func String() Decoder {
	return func(b []byte, i interface{}) error {
		_, ok := i.(*string)
		if !ok {
			return ErrWrongType{expected: "string"}
		}
		if err := json.Unmarshal(b, i); err != nil {
			return err
		}
		return nil
	}
}
