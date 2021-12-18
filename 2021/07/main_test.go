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
	require.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, input)
}

func TestCalculateOptimalHorizontalPosition(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, input)

	require.Equal(t, 2, calculateOptimalHorizontalPostion(input))
}
