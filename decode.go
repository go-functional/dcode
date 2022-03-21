package dcode

// Decode decodes b into i using d
func Decode[T any](b []byte, d Decoder[T]) (T, error) {
	var zero T
	ret, err := d.Decode(b)
	if err != nil {
		return zero, err
	}
	return ret, nil
}
