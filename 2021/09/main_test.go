package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte
var expectedCave = Cave{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}

func TestReadInput(t *testing.T) {
	cave, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedCave, cave)
}

func TestIsLowPoint(t *testing.T) {
	type test struct {
		position Point
		expected bool
	}

	tests := []test{
		{position: Point{X: 0, Y: 0}, expected: false},
		{position: Point{X: 1, Y: 0}, expected: true},
		{position: Point{X: 1, Y: 1}, expected: false},
		{position: Point{X: 9, Y: 0}, expected: true},
		{position: Point{X: 6, Y: 4}, expected: true},
	}

	for _, tc := range tests {
		got := isLowPoint(expectedCave, tc.position)
		require.Equal(t, tc.expected, got)
	}
}

func TestFindRiskLevelSum(t *testing.T) {
	sum := findRiskLevelSum(expectedCave)
	require.Equal(t, 15, sum)
}

func TestFindThreeLargestBasinMultiple(t *testing.T) {
	m := findThreeLargestBasinMultiple(expectedCave)
	require.Equal(t, 1134, m)
}

func TestFindBasinSize(t *testing.T) {
	type test struct {
		position Point
		expected int
	}

	tests := []test{
		{position: Point{X: 1, Y: 0}, expected: 3},
		{position: Point{X: 9, Y: 0}, expected: 9},
		{position: Point{X: 2, Y: 2}, expected: 14},
		{position: Point{X: 6, Y: 4}, expected: 9},
	}

	for _, tc := range tests {
		basin := map[Point]bool{tc.position: true}
		findBasinSize(expectedCave, tc.position, basin)
		require.Equal(t, tc.expected, len(basin))
	}
}
