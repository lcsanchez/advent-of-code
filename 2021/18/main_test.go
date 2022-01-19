package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestReadInput(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Len(t, input, 2)
	require.Equal(t, &Node{
		left: &Node{
			left: &Node{
				left: &Node{
					left:  &Node{data: 4},
					right: &Node{data: 3},
				},
				right: &Node{data: 4},
			},
			right: &Node{data: 4},
		},
		right: &Node{
			left: &Node{data: 7},
			right: &Node{
				left: &Node{
					left:  &Node{data: 8},
					right: &Node{data: 4},
				},
				right: &Node{data: 9},
			},
		},
	}, input[0])

	require.Equal(t, &Node{
		left:  &Node{data: 1},
		right: &Node{data: 1},
	}, input[1])
}
