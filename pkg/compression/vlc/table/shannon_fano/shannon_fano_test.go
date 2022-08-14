package shannon_fano

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBestDividerPosition(t *testing.T) {
	tests := []struct {
		name     string
		codes    []code
		expected int
	}{
		{
			"one element",
			[]code{
				{Quantity: 2},
			},
			0,
		},
		{
			"two elements",
			[]code{
				{Quantity: 2},
				{Quantity: 2},
			},
			1,
		},
		{
			"three elements",
			[]code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			1,
		},
		{
			"many elements",
			[]code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			2,
		},
		{
			"uncertainty (need rightmost)",
			[]code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			1,
		},
		{
			"uncertainty (need rightmost)",
			[]code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, bestDividerPosition(tt.codes))
		})
	}
}

func TestAssignCodes(t *testing.T) {
	tests := []struct {
		name     string
		codes    []code
		expected []code
	}{
		{
			"base test",
			[]code{
				{Quantity: 2},
				{Quantity: 2},
			},
			[]code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			"three items test, centrain position",
			[]code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			[]code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
		{
			"three items test, uncentrain position",
			[]code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			[]code{
				{Quantity: 1, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)
			require.Equal(t, tt.expected, tt.codes)
		})
	}
}

func TestBuild(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected encodingTable
	}{
		{
			"base test",
			"abbbcc",
			encodingTable{
				'a': code{Quantity: 1, Char: 'a', Bits: 3, Size: 2},
				'b': code{Quantity: 3, Char: 'b', Bits: 0, Size: 1},
				'c': code{Quantity: 2, Char: 'c', Bits: 2, Size: 2},
			},
		},
		{
			"base test (uncentrain)",
			"aabb",
			encodingTable{
				'a': code{Quantity: 2, Char: 'a', Bits: 0, Size: 1},
				'b': code{Quantity: 2, Char: 'b', Bits: 1, Size: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, build(newChatStat(tt.str)))
		})
	}
}
