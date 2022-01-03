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
	Template: "NNCB",
	Rules: []*Rule{
		{Match: "CH", Insert: "B"},
		{Match: "HH", Insert: "N"},
		{Match: "CB", Insert: "H"},
		{Match: "NH", Insert: "C"},
		{Match: "HB", Insert: "C"},
		{Match: "HC", Insert: "B"},
		{Match: "HN", Insert: "C"},
		{Match: "NN", Insert: "C"},
		{Match: "BH", Insert: "H"},
		{Match: "NC", Insert: "B"},
		{Match: "NB", Insert: "B"},
		{Match: "BN", Insert: "B"},
		{Match: "BB", Insert: "N"},
		{Match: "BC", Insert: "B"},
		{Match: "CC", Insert: "N"},
		{Match: "CN", Insert: "C"},
	},
}

func TestReadInput(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedInput, input)
}
