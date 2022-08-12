package vlc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSplitByChunks(t *testing.T) {
	tests := []struct {
		str         string
		chunkLength int
		expected    BinaryChunks
	}{
		{"011101011110101000001100", 8, BinaryChunks{
			"01110101",
			"11101010",
			"00001100",
		}},
		{"0111010101", 4, BinaryChunks{
			"0111",
			"0101",
			"0100",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			require.Equal(t, tt.expected, splitByChunks(tt.str, tt.chunkLength))
		})
	}
}

func TestBinaryChunksJoin(t *testing.T) {
	tests := []struct {
		name     string
		bcs      BinaryChunks
		expected string
	}{
		{"empty", BinaryChunks{}, ""},
		{"empty", BinaryChunks{BinaryChunk("0000"), BinaryChunk("1111")}, "00001111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.bcs.Join())
		})
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected BinaryChunks
	}{
		{"base test", []byte{20, 30, 60, 18}, BinaryChunks{"00010100", "00011110", "00111100", "00010010"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, NewBinChunks(tt.data))
		})
	}
}
