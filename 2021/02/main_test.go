package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

func TestNavigate(t *testing.T) {
	s := &Submarine{Position: &Point{}}
	err := navigateSubmarine(s, bytes.NewReader(testInput), SubmarineCommandBuilder)
	require.NoError(t, err)

	want := &Submarine{
		Position: &Point{X: 15, Y: 10},
	}

	if diff := cmp.Diff(want, s); diff != "" {
		t.Errorf("readPoints() mismatch (-want +got):\n%s", diff)
	}
}

func TestNavigateSubmarineWithAim(t *testing.T) {
	s := &Submarine{Position: &Point{}}
	err := navigateSubmarine(s, bytes.NewReader(testInput), SubmarineWithAimCommandBuilder)
	require.NoError(t, err)

	want := &Submarine{
		Aim:      10,
		Position: &Point{X: 15, Y: 60},
	}

	if diff := cmp.Diff(want, s); diff != "" {
		t.Errorf("readPoints() mismatch (-want +got):\n%s", diff)
	}
}
