package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const chunkLength = 8

type BinaryChunk string

type BinaryChunks []BinaryChunk

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, x := range data {
		res = append(res, NewBinChunk(x))
	}

	return res
}

func NewBinChunk(b byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", b))
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}

	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunkLength)

	if err != nil {
		panic("cant convert binary chunk to byte: " + err.Error())
	}

	return byte(num)
}

// Join join chunks to string without spaces
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}

	return buf.String()
}

// splitByChunks splits string by chunks
func splitByChunks(str string, chunkLength int) BinaryChunks {
	chunkCount := utf8.RuneCountInString(str) / chunkLength
	if len(str)%chunkLength != 0 {
		chunkCount++
	}

	bc := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder

	for _, x := range str {
		buf.WriteRune(x)
		if buf.Len()%chunkLength == 0 {
			bc = append(bc, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		zeros := strings.Repeat("0", chunkLength-buf.Len())
		buf.WriteString(zeros)
		bc = append(bc, BinaryChunk(buf.String()))
	}

	return bc
}
