package dcode

func (t *TheSuite) TestSimpleDecode() {
	r := t.Require()
	jsonBytes := []byte(`{"a":1}`)
	var i int
	r.NoError(Decode(jsonBytes, Field("a", Int()), &i))
	r.Equal(1, i)
}
