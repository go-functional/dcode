package dcode

func (t *TheSuite) TestSingleField() {
	r := t.Require()
	b := []byte(jsonStr)
	decoder := Field("simple", Int())
	actual, err := decoder.call(b)
	r.NoError(err)
	r.Equal(1, actual)
}
