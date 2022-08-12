package compression

type Encoder interface {
	Encode(str string) []byte
	Extension() string
}

type Decoder interface {
	Decode(data []byte) string
}
