package dcode

import (
	tree "github.com/bmatsuo/go-jsontree"
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
	return newDecoder(func(t *tree.JsonTree) (interface{}, error) {
		// First check for a leaf in the JSON tree.
		// In other words, this is the very inner Field() call, or
		// the decoder passed to Into()
		i, err := decoder.call(t)
		if err == nil {
			return i, nil
		}

		// Otherwise, if we aren't a leaf node, then try to get
		// a subtree, and then call the decoder on that subtree.
		//
		// In other words, traverse the tree
		subTree := t.Get(name)
		if subTree.Err() != nil {
			return nil, subTree.Err()
		}
		return decoder.call(subTree)
	})
}
