package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInputBytes []byte
var testInput = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestReadNums(t *testing.T) {
	got, err := readNums(bytes.NewReader(testInputBytes))
	require.NoError(t, err)

	want := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	require.Equal(t, want, got)

}

func TestCountDepthIncrease(t *testing.T) {
	count := countDepthIncrease(testInput)

	require.Equal(t, 7, count)
}

func TestMakeGroups(t *testing.T) {
	got := makeGroups(testInput, 3)

	want := []int{
		607,
		618,
		618,
		617,
		647,
		716,
		769,
		792,
	}

	require.Equal(t, want, got)
}

func TestCountDepthGroupIncrease(t *testing.T) {
	count := countDepthGroupIncrease(testInput, 3)

	require.Equal(t, 5, count)
}
