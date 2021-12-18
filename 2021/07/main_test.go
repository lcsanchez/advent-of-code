package main

import (
	"bytes"
	_ "embed"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestReadInput(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)

	expected := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	sort.Ints(expected)

	require.Equal(t, expected, input)
}

func TestCalculateOptimalHorizontalPositionPartOne(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)

	position, cost := calculateOptimalHorizontalPostion(input, identity)
	require.Equal(t, 2, position)
	require.Equal(t, 37, cost)
}

func TestCalculateOptimalHorizontalPositionPartTwo(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)

	position, cost := calculateOptimalHorizontalPostion(input, sumToN)
	require.Equal(t, 5, position)
	require.Equal(t, 168, cost)
}
