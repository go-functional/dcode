package dcode

import (
	tree "github.com/bmatsuo/go-jsontree"
)

// Float64 decodes any JSON field into a float64
func Float64() Decoder {
	return newDecoder(func(t *tree.JsonTree) (interface{}, error) {
		ret, err := t.Number()
		if err != nil {
			return nil, err
		}
		return ret, nil
	})
}
