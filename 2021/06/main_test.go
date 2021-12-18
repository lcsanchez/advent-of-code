package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

//go:embed testdata/test-input-expected.txt
var testInputExpected []byte

func TestReadInput(t *testing.T) {
	fish, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, fish[0])
}

func TestCountTotalFish(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Len(t, input, 1)

	require.Equal(t, 5, countTotalFish(input[0]))
}

func TestCalculateFish(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Len(t, input, 1)

	actualFish := input[0]

	expectedFish, err := readInput(bytes.NewReader(testInputExpected))
	require.NoError(t, err)

	for i, expected := range expectedFish {
		calculateFishAfterOneDay(actualFish)
		require.Equal(t, expected, actualFish, "day %d", i+1)
	}
}
