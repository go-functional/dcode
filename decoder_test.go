package dcode

import "fmt"

func (t *TheSuite) SimpleInt() {
	r := t.Require()
	ints := []int{1, 2, 3, 4, 5, 1000}
	for _, i := range ints {
		ret, err := DecodeString(Int(), fmt.Sprintf("%d", i))
		r.NoError(err)
		r.Equal(1, ret)
	}

	notInts := []string{`"abc"`, `"dev"`}
	for _, i := range notInts {
		ret, err := DecodeString(Int(), i)
		r.True(err != nil)
		r.Nil(ret)
	}
}

func (t *TheSuite) SimpleString() {
	r := t.Require()
	for _, i := range []string{`"this is a thing"`, `"this is another thing"`} {
		ret, err := DecodeString(String(), i)
		r.NoError(err)
		r.Equal(i, ret)
	}
}
