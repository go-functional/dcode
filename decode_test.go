package dcode

func (t *TheSuite) TestSimpleDecode() {
	r := t.Require()
	jsonBytes := []byte(jsonStr)
	var i int
	r.NoError(Decode(jsonBytes, Field("simple", Int()), &i))
	r.Equal(1, i)
}
