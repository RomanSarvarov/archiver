package vlc

import (
	"strings"
	"unicode"
)

type EncoderDecoder struct{}

func NewEncoderDecoder() EncoderDecoder {
	return EncoderDecoder{}
}

type encodingTable map[rune]string

var table = getEncodingTable()

func (ed EncoderDecoder) Encode(str string) []byte {
	str = prepareText(str)

	str = encodeBin(str)

	chunks := splitByChunks(str, chunkLength)

	return chunks.Bytes()
}

func (ed EncoderDecoder) Extension() string {
	return "vlc"
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	str := NewBinChunks(encodedData).Join()

	dt := getEncodingTable().DecodingTree()

	return exportText(dt.Decode(str))
}

// prepareText format text
// capitalized to lower cased with "!" before
func prepareText(str string) string {
	if str == "" {
		return ""
	}

	var buf strings.Builder

	for _, x := range str {
		if unicode.IsUpper(x) {
			buf.WriteString("!")
			buf.WriteRune(unicode.ToLower(x))
		} else {
			buf.WriteRune(x)
		}
	}

	return buf.String()
}

// exportText format text
// lower to capitalized if "!" before
func exportText(str string) string {
	if str == "" {
		return ""
	}

	var buf strings.Builder

	nextLetterShouldBeCapitalized := false

	for _, x := range str {
		if x == '!' {
			nextLetterShouldBeCapitalized = true
			continue
		}

		if nextLetterShouldBeCapitalized && !unicode.IsLower(x) {
			buf.WriteRune('!')
			nextLetterShouldBeCapitalized = false
		}

		if nextLetterShouldBeCapitalized && unicode.IsLower(x) {
			buf.WriteRune(unicode.ToUpper(x))
		} else {
			buf.WriteRune(x)
		}

		nextLetterShouldBeCapitalized = false
	}

	return buf.String()
}

// encodeBin converts string to binary
func encodeBin(str string) string {
	var buf strings.Builder

	for _, x := range str {
		buf.WriteString(bin(x))
	}

	return buf.String()
}

func bin(r rune) string {
	b, ok := table[r]

	if !ok {
		panic("unknown character: " + string(r))
	}

	return b
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}
