package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestFindLowestRiskPath(t *testing.T) {
	cave, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)

	require.Equal(t, 40, findLowestRiskPath(cave))

	largeCave := buildLargeCave(cave, 5)
	require.Equal(t, 315, findLowestRiskPath(largeCave))
}
