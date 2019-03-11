package dcode

import (
	"encoding/json"
	"fmt"
)

// Field builds up a traversal. Here's an example that starts with
// "field1", then goes down to "field2", and then into "field3", and decodes
// that value into an int:
//
//	dcoder := Field("field1", Field("field2", Field("field3", Int())))
//
// If you have the following JSON:
//
//	{
//		"field1": {
//			"field2": {
//				"field3": 123
//			}
//		}
//	}
//
// Then the following code would decode the integer 123 into the variable i
//
//	var i int
//	Decode(dcoder, jsonBytes, &i) // make sure to check the error returned!
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
