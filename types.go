package dcode

import "fmt"

type ErrWrongType struct {
	expected string
	actual   interface{}
}

func (e ErrWrongType) Error() string {
	return fmt.Sprintf("Expected a %s, got a %T", e.expected, e.actual)
}
