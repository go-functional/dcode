package dcode

// Builder is a convenience struct that lets you write decoders as a chain,
// rather than in the "pure functional" style. Instead of this:
//
//	dcoder := Field("first", Field("second", Field("third", Int())))
//
// You can write this:
//
//	dcoder := First("first").Then("second").Then("third").Into(Int())
type Builder struct {
	fields []string
}

// Then traverses one level down into the JSON object. See the documentation
// in Builder for a complete explanation
func (b Builder) Then(field string) Builder {
	return Builder{fields: append(b.fields, field)}
}

// Into returns a Decoder that decodes the value at the current travere level
// using d. See the documentation in Builder for a complete explanation
func (b Builder) Into(d Decoder) Decoder {
	// build up the decoder from the very last field (i.e. depth first)
	//
	// TODO: consider doing this recursively
	decoder := d
	for i := len(b.fields) - 1; i >= 0; i-- {
		field := b.fields[i]
		decoder = Field(field, decoder)
	}
	return decoder
}

// First returns a Builder that starts at field in the JSON object
func First(field string) Builder {
	return Builder{fields: []string{field}}
}
