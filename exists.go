package dcode

func Exists(decoder Decoder, bytes []byte) bool {
	if _, err := decoder.call(bytes); err != nil {
		return false
	}
	return true
}
