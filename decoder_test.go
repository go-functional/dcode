package dcode

import "fmt"

func (t *TheSuite) TestSimpleInt() {
	r := t.Require()
	ints := []int{1, 2, 3, 4, 5, 1000}
	for i, expected := range ints {
		var actual int
		r.NoError(
			Int()([]byte(fmt.Sprintf("%d", expected)), &actual),
			"for iteration %d, int %d",
			i,
			expected,
		)
		r.Equal(expected, actual, "expected int %d", actual)
	}

	notInts := []string{`"abc"`, `"dev"`}
	for _, notInt := range notInts {
		var actual int
		r.True(
			Int()([]byte(notInt), &actual) != nil,
		)
		r.Equal(0, actual)
	}
}

func (t *TheSuite) TestSimpleString() {
	r := t.Require()
	strings := []string{
		`"this is a thing"`,
		`"this is another thing"`,
	}
	for i, expected := range strings {
		b := []byte(expected)
		var actual string
		r.NoError(String()(b, &actual), "for iteration %d", i)
		r.Equal(expected, `"`+actual+`"`, "for iteration %d", i)
	}
}
