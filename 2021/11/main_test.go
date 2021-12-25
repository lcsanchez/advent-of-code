package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte
var expected = [][]int{
	{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
	{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
	{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
	{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
	{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
	{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
	{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
	{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
	{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
	{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
}

type ExpectedStep struct {
	Flashes int
	Grid    [][]int
}

var expectedSteps = []*ExpectedStep{
	{
		Grid: [][]int{
			{6, 5, 9, 4, 2, 5, 4, 3, 3, 4},
			{3, 8, 5, 6, 9, 6, 5, 8, 2, 2},
			{6, 3, 7, 5, 6, 6, 7, 2, 8, 4},
			{7, 2, 5, 2, 4, 4, 7, 2, 5, 7},
			{7, 4, 6, 8, 4, 9, 6, 5, 8, 9},
			{5, 2, 7, 8, 6, 3, 5, 7, 5, 6},
			{3, 2, 8, 7, 9, 5, 2, 8, 3, 2},
			{7, 9, 9, 3, 9, 9, 2, 2, 4, 5},
			{5, 9, 5, 7, 9, 5, 9, 6, 6, 5},
			{6, 3, 9, 4, 8, 6, 2, 6, 3, 7},
		},
	},
	{
		Flashes: 35,
		Grid: [][]int{
			{8, 8, 0, 7, 4, 7, 6, 5, 5, 5},
			{5, 0, 8, 9, 0, 8, 7, 0, 5, 4},
			{8, 5, 9, 7, 8, 8, 9, 6, 0, 8},
			{8, 4, 8, 5, 7, 6, 9, 6, 0, 0},
			{8, 7, 0, 0, 9, 0, 8, 8, 0, 0},
			{6, 6, 0, 0, 0, 8, 8, 9, 8, 9},
			{6, 8, 0, 0, 0, 0, 5, 9, 4, 3},
			{0, 0, 0, 0, 0, 0, 7, 4, 5, 6},
			{9, 0, 0, 0, 0, 0, 0, 8, 7, 6},
			{8, 7, 0, 0, 0, 0, 6, 8, 4, 8},
		},
	},
	{
		Flashes: 45,
		Grid: [][]int{
			{0, 0, 5, 0, 9, 0, 0, 8, 6, 6},
			{8, 5, 0, 0, 8, 0, 0, 5, 7, 5},
			{9, 9, 0, 0, 0, 0, 0, 0, 3, 9},
			{9, 7, 0, 0, 0, 0, 0, 0, 4, 1},
			{9, 9, 3, 5, 0, 8, 0, 0, 6, 3},
			{7, 7, 1, 2, 3, 0, 0, 0, 0, 0},
			{7, 9, 1, 1, 2, 5, 0, 0, 0, 9},
			{2, 2, 1, 1, 1, 3, 0, 0, 0, 0},
			{0, 4, 2, 1, 1, 2, 5, 0, 0, 0},
			{0, 0, 2, 1, 1, 1, 9, 0, 0, 0},
		},
	},
}

func TestReadInput(t *testing.T) {
	cave, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expected, cave)
}

func TestStep(t *testing.T) {
	grid := copyGrid(expected)

	for i := 0; i < len(expectedSteps); i++ {
		flashes := step(grid)
		require.Equal(t, expectedSteps[i].Grid, grid, "step %d", i+1)
		require.Equal(t, expectedSteps[i].Flashes, flashes, "step %d", i+1)
	}
}
