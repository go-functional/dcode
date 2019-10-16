package dcode

import (
	tree "github.com/bmatsuo/go-jsontree"
)

// Bool decodes any JSON field into a float64
func Bool() Decoder {
	return newDecoder(func(t *tree.JsonTree) (interface{}, error) {
		ret, err := t.Boolean()
		if err != nil {
			var zero bool
			return zero, err
		}
		return ret, nil
	})
}
