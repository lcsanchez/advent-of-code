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
	require.Len(t, input, 5)
}

func TestSolve(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	s := solve(input)

	require.Len(t, s.beacons, 79)
}

func TestManhattanDistance(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	s := solve(input)

	require.Equal(t, 3621, manhattanDistance(s))
}
