package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

var expectedInput = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestReadInput(t *testing.T) {
	got, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedInput, got)
}
