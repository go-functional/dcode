package dcode

import (
	tree "github.com/bmatsuo/go-jsontree"
)

// Int decodes any JSON field into an int
func Int() Decoder {
	return newDecoder(func(t *tree.JsonTree) (interface{}, error) {
		ret, err := t.Number()
		if err != nil {
			return nil, err
		}
		return int(ret), nil
	})
}
