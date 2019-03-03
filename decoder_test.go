package dcode

import "fmt"

func (t *TheSuite) TestSimpleInt() {
	r := t.Require()
	ints := []int{1, 2, 3, 4, 5, 1000}
	for _, i := range ints {
		ret, err := DecodeString(Int(), fmt.Sprintf("%d", i))
		r.NoError(err, "for int %d", i)
		r.Equal(i, ret, "expected int %d", i)
	}

	notInts := []string{`"abc"`, `"dev"`}
	for _, i := range notInts {
		ret, err := DecodeString(Int(), i)
		r.True(err != nil)
		r.Nil(ret)
	}
}

func (t *TheSuite) TestSimpleString() {
	r := t.Require()
	for _, i := range []string{`"this is a thing"`, `"this is another thing"`} {
		ret, err := DecodeString(String(), i)
		r.NoError(err, "for string %s", i)
		r.Equal(i, fmt.Sprintf(`"%s"`, ret), ret, "for string %s", i)
	}
}
