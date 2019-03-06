package dcode

func (t *TheSuite) TestSingleField() {
	r := t.Require()
	var i int
	b := []byte(`{"a": 1}`)
	d := Field("a", Int())
	r.NoError(Decode(d, b, &i))
	r.Equal(1, i)
}
