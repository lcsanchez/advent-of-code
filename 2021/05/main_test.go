package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestReadPoint(t *testing.T) {
	point, err := readPoint("0,9")
	require.NoError(t, err)
	require.Equal(t, &Point{X: 0, Y: 9}, point)
}

func TestReadInput(t *testing.T) {
	lines, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, []*Line{
		{Start: &Point{X: 0, Y: 9}, End: &Point{X: 5, Y: 9}},
		{Start: &Point{X: 8, Y: 0}, End: &Point{X: 0, Y: 8}},
		{Start: &Point{X: 9, Y: 4}, End: &Point{X: 3, Y: 4}},
		{Start: &Point{X: 2, Y: 2}, End: &Point{X: 2, Y: 1}},
		{Start: &Point{X: 7, Y: 0}, End: &Point{X: 7, Y: 4}},
		{Start: &Point{X: 6, Y: 4}, End: &Point{X: 2, Y: 0}},
		{Start: &Point{X: 0, Y: 9}, End: &Point{X: 2, Y: 9}},
		{Start: &Point{X: 3, Y: 4}, End: &Point{X: 1, Y: 4}},
		{Start: &Point{X: 0, Y: 0}, End: &Point{X: 8, Y: 8}},
		{Start: &Point{X: 5, Y: 5}, End: &Point{X: 8, Y: 2}},
	}, lines)
}
