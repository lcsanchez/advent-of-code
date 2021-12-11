package main

import (
	"bytes"
	_ "embed"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestReadReport(t *testing.T) {
	report, err := readReport(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}, report)
}

func TestCalculatePowerConsumption(t *testing.T) {
	powerConsumption, err := calculatePowerConsumption(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, 198, powerConsumption)

}

func TestCalculateOxygenRating(t *testing.T) {
	report, err := readReport(bytes.NewReader(testInput))
	require.NoError(t, err)
	sort.Strings(report)

	rating, err := calculateRating(report, 0, O2RatingFilterFunc)
	require.NoError(t, err)
	require.Equal(t, 23, rating)
}

func TestCalculateC02ScrubberRating(t *testing.T) {
	report, err := readReport(bytes.NewReader(testInput))
	require.NoError(t, err)
	sort.Strings(report)

	rating, err := calculateRating(report, 0, CO2RatingFilterFunc)
	require.NoError(t, err)
	require.Equal(t, 10, rating)
}

func TestCalculateLifeSupportRating(t *testing.T) {
	rating, err := calculateLifeSupportRating(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, 230, rating)
}
