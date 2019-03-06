package dcode

func (t *TheSuite) TestBasicBuilder() {
	r := t.Require()

	decoder := First("a").Then("otherb").Then("otherc").Into(Int())
	var i int
	jsonBytes := []byte(jsonStr)
	r.NoError(Decode(jsonBytes, decoder, &i))
	r.Equal(200, i)
}
