package dcode

func Exists(decoder Decoder, bytes []byte) bool {
	var i interface{}
	if err := Decode(decoder, bytes, i); err != nil {
		return false
	}
	return true
}
