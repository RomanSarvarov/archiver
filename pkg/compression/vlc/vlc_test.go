package vlc

import (
	"testing"
)
import "github.com/stretchr/testify/require"

func TestPrepareText(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{str: "", expected: ""},
		{str: "abc", expected: "abc"},
		{str: "aBc", expected: "a!bc"},
		{str: "123", expected: "123"},
		{str: "Hello World!", expected: "!hello !world!"},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			require.Equal(t, tt.expected, prepareText(tt.str))
		})
	}
}

func TestEncodeBin(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{"!roman !sarvarov", "0010000100010001000011011100001100100001010110100000000001011010001000100000001"},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			require.Equal(t, tt.expected, encodeBin(tt.str))
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []byte
	}{
		{"base test", "Hello! My name is Roman", []byte{32, 233, 36, 196, 140, 128, 192, 240, 97, 221, 43, 144, 136, 134, 224}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewEncoderDecoder()
			require.Equal(t, tt.expected, encoder.Encode(tt.str))
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		encoded []byte
		decoded string
	}{
		{"base test", []byte{32, 233, 36, 196, 140, 128, 192, 240, 97, 221, 43, 144, 136, 134, 224}, "Hello! My name is Roman"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := NewEncoderDecoder()
			require.Equal(t, tt.decoded, decoder.Decode(tt.encoded))
		})
	}
}
