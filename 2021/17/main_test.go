package main

import (
	"bytes"
	_ "embed"
	"math"
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

func TestFindSummationInRange(t *testing.T) {
	i, sum := findSummationInRange(&Tuple{First: 20, Second: 30})
	require.Equal(t, 6, i)
	require.Equal(t, 21, sum)
}

func TestCalculateSummation(t *testing.T) {
	sum := calculateSummation(9)
	require.Equal(t, 45, sum)
}

func TestCountStartingVelocities(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, 112, countStartingVelocities(input))
}

func TestCalculateYVelocities(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, 3, calculateYVelocitiesInSteps(input.Y, 2, 2))
	require.Equal(t, 3, calculateYVelocitiesInSteps(input.Y, 4, math.MaxInt64))
}
