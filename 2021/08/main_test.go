package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestReadInput(t *testing.T) {
	input, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)

	expected := []*Entry{
		{
			SignalPatterns: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
			OutputValues:   []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
		},
		{
			SignalPatterns: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"},
			OutputValues:   []string{"fcgedb", "cgb", "dgebacf", "gc"},
		},
		{
			SignalPatterns: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"},
			OutputValues:   []string{"cg", "cg", "fdcagb", "cbg"},
		},
		{
			SignalPatterns: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"},
			OutputValues:   []string{"efabcd", "cedba", "gadfec", "cb"},
		},
		{
			SignalPatterns: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"},
			OutputValues:   []string{"gecf", "egdcabf", "bgf", "bfgea"},
		},
		{
			SignalPatterns: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"},
			OutputValues:   []string{"gebdcfa", "ecba", "ca", "fadegcb"},
		},
		{
			SignalPatterns: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"},
			OutputValues:   []string{"cefg", "dcbef", "fcge", "gbcadfe"},
		},
		{
			SignalPatterns: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"},
			OutputValues:   []string{"ed", "bcgafe", "cdgba", "cbgef"},
		},
		{
			SignalPatterns: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"},
			OutputValues:   []string{"gbdfcae", "bgc", "cg", "cgb"},
		},
		{
			SignalPatterns: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"},
			OutputValues:   []string{"fgae", "cfgab", "fg", "bagce"},
		},
	}

	require.Equal(t, expected, input)
}
