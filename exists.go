package dcode

// Exists checks whether the value that decoder points to (in the JSON bytes)
// exists and is the expected type
func Exists(decoder Decoder, bytes []byte) bool {
	if _, err := decoder.call(bytes); err != nil {
		return false
	}
	return true
}
