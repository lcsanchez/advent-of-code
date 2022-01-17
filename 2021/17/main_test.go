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
	require.Equal(t, &Input{
		X: &Tuple{
			First:  20,
			Second: 30,
		},
		Y: &Tuple{
			First:  -5,
			Second: -10,
		},
	}, input)
}
