package dcode

import (
	"fmt"
	"strconv"
)

func (t *TheSuite) TestSimpleInt() {
	r := t.Require()
	ints := []int{1, 2, 3, 4, 5, 1000}
	decoder := Int()
	for i, expected := range ints {
		expectedBytes := []byte(strconv.Itoa(expected))
		actual, err := decoder.call(expectedBytes)
		r.NoError(err, "for iteration %d, int %d", i, expected)
		r.Equal(expected, actual, "expected int %d", actual)
	}

	notInts := []string{`"abc"`, `"dev"`}
	for _, notInt := range notInts {
		actual, err := decoder.call([]byte(notInt))
		r.True(err != nil)
		r.Equal(0, actual)
	}
}

func (t *TheSuite) TestSimpleString() {
	r := t.Require()
	decoder := String()
	strings := []string{
		`"this is a thing"`,
		`"this is another thing"`,
	}
	for i, expected := range strings {
		b := []byte(expected)
		actual, err := decoder.call(b)
		r.NoError(err, "for iteration %d", i)
		actualJSONStr := fmt.Sprintf(`"%s"`, actual)
		r.Equal(expected, actualJSONStr, "for iteration %d", i)
	}
}
