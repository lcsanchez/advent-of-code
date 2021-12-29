package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

var expectedPaths = []*Path{
	{a: "start", b: "A"},
	{a: "start", b: "b"},
	{a: "A", b: "c"},
	{a: "A", b: "b"},
	{a: "b", b: "d"},
	{a: "A", b: "end"},
	{a: "b", b: "end"},
}

func TestReadInput(t *testing.T) {
	paths, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedPaths, paths)
}
