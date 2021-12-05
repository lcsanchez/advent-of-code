package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestReadPoints(t *testing.T) {
	input := strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2
`)

	want := []*Point{
		{X: 5},
		{Y: 5},
		{X: 8},
		{Y: -3},
		{Y: 8},
		{X: 2},
	}

	got, err := readPoints(input)
	require.NoError(t, err)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("readPoints() mismatch (-want +got):\n%s", diff)
	}
}

func TestNavigate(t *testing.T) {
	course := []*Point{
		{X: 5},
		{Y: 5},
		{X: 8},
		{Y: -3},
		{Y: 8},
		{X: 2},
	}

	got := navigate(course)

	if diff := cmp.Diff(&Point{X: 15, Y: 10}, got); diff != "" {
		t.Errorf("readPoints() mismatch (-want +got):\n%s", diff)
	}
}
