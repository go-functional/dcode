package dcode

type Builder struct {
	fields []string
}

func (b Builder) Then(field string) Builder {
	return Builder{fields: append(b.fields, field)}
}

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

func First(field string) Builder {
	return Builder{fields: []string{field}}
}
