package dcode

import (
	"encoding/json"
)

func Int() Decoder {
	return func(b []byte, i interface{}) error {
		_, ok := i.(*int)
		if !ok {
			return ErrWrongType{expected: "int"}
		}
		if err := json.Unmarshal(b, i); err != nil {
			return err
		}
		return nil
	}
}
