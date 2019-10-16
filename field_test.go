package dcode

func (t *TheSuite) TestSingleField() {
	r := t.Require()
	b := []byte(jsonStr)
	tree, err := t.getTree(b)
	r.NoError(err)
	decoder := Field("simple", Int())
	actual, err := decoder.call(tree)
	r.NoError(err)
	r.Equal(1, actual)
}
