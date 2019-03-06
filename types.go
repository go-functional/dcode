package dcode

import "fmt"

type ErrWrongType struct {
	expected string
}

func (e ErrWrongType) Error() string {
	return fmt.Sprintf("expected a %s", e.expected)
}
