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
		{Match: "CH", Insert: 'B'},
		{Match: "HH", Insert: 'N'},
		{Match: "CB", Insert: 'H'},
		{Match: "NH", Insert: 'C'},
		{Match: "HB", Insert: 'C'},
		{Match: "HC", Insert: 'B'},
		{Match: "HN", Insert: 'C'},
		{Match: "NN", Insert: 'C'},
		{Match: "BH", Insert: 'H'},
		{Match: "NC", Insert: 'B'},
		{Match: "NB", Insert: 'B'},
		{Match: "BN", Insert: 'B'},
		{Match: "BB", Insert: 'N'},
		{Match: "BC", Insert: 'B'},
		{Match: "CC", Insert: 'N'},
		{Match: "CN", Insert: 'C'},
	},
}

func TestReadInput(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedInput, input)
}

func TestApplyPolymerNTimes(t *testing.T) {
	type test struct {
		input string
		times int
		min   int
		max   int
	}

	tests := []test{
		{input: "NNCB", times: 1, min: 1, max: 2},
		{input: "NNCB", times: 2, min: 1, max: 6},
		{input: "NBBBCNCCNBBNBNBBCHBHHBCHB", times: 2, min: 1, max: 13},
	}

	for _, tc := range tests {
		min, max := applyPolymer(tc.input, expectedInput.Rules, tc.times)
		require.Equal(t, tc.min, min)
		require.Equal(t, tc.max, max)
	}
}
