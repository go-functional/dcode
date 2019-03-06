package dcode

func DecodeString(d Decoder, s string, iface interface{}) error {
	return Decode(d, []byte(s), iface)
}

func Decode(d Decoder, b []byte, iface interface{}) error {
	return d(b, iface)
}
