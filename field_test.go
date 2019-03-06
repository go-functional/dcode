package dcode

func (t *TheSuite) TestSingleField() {
	r := t.Require()
	b := []byte(`{"a": 1}`)
	decoder := Field("a", Int())
	actual, err := decoder(b)
	r.NoError(err)
	r.Equal(1, actual)
}
