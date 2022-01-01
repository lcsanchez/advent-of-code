package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

var expectedInput = &Input{
	Points: []Point{
		{X: 6, Y: 10},
		{X: 0, Y: 14},
		{X: 9, Y: 10},
		{X: 0, Y: 3},
		{X: 10, Y: 4},
		{X: 4, Y: 11},
		{X: 6, Y: 0},
		{X: 6, Y: 12},
		{X: 4, Y: 1},
		{X: 0, Y: 13},
		{X: 10, Y: 12},
		{X: 3, Y: 4},
		{X: 3, Y: 0},
		{X: 8, Y: 4},
		{X: 1, Y: 10},
		{X: 2, Y: 14},
		{X: 8, Y: 10},
		{X: 9, Y: 0},
	},
	Folds: []Fold{
		{FoldAxis: FoldAxisY, Value: 7},
		{FoldAxis: FoldAxisX, Value: 5},
	},
}

func TestReadInput(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedInput, input)
}

func TestFold(t *testing.T) {
	expected := [][]Point{
		{
			{X: 6, Y: 4},
			{X: 0, Y: 0},
			{X: 9, Y: 4},
			{X: 0, Y: 3},
			{X: 10, Y: 4},
			{X: 4, Y: 3},
			{X: 6, Y: 0},
			{X: 6, Y: 2},
			{X: 4, Y: 1},
			{X: 0, Y: 1},
			{X: 10, Y: 2},
			{X: 3, Y: 4},
			{X: 3, Y: 0},
			{X: 8, Y: 4},
			{X: 1, Y: 4},
			{X: 2, Y: 0},
			// duplicate {X: 8, Y: 4},
			{X: 9, Y: 0},
		},
		{
			{X: 4, Y: 4},
			{X: 0, Y: 0},
			{X: 1, Y: 4},
			{X: 0, Y: 3},
			{X: 0, Y: 4},
			{X: 4, Y: 3},
			{X: 4, Y: 0},
			{X: 4, Y: 2},
			{X: 4, Y: 1},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 3, Y: 4},
			{X: 3, Y: 0},
			{X: 2, Y: 4},
			// duplicate {X: 1, Y: 4},
			{X: 2, Y: 0},
			{X: 1, Y: 0},
		},
	}
	points := expectedInput.Points
	for i, f := range expectedInput.Folds {
		points = fold(points, f)
		require.ElementsMatch(t, expected[i], points)
	}
}
