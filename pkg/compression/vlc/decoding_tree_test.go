package vlc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodingTableDecodingTree(t *testing.T) {
	tests := []struct {
		name     string
		et       encodingTable
		expected DecodingTree
	}{
		{"base tree test", encodingTable{'a': "110", 'b': "0100", 'c': "11110"}, DecodingTree{
			One: &DecodingTree{
				One: &DecodingTree{
					One: &DecodingTree{
						One: &DecodingTree{
							Zero: &DecodingTree{
								Value: "c",
							},
						},
					},
					Zero: &DecodingTree{
						Value: "a",
					},
				},
			},
			Zero: &DecodingTree{
				One: &DecodingTree{
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							Value: "b",
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.et.DecodingTree())

		})
	}
}
