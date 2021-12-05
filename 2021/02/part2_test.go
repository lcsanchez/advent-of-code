package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestNavigateSubmarine(t *testing.T) {
	input := strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2
`)

	commands, err := readCommands(input)
	require.NoError(t, err)

	s := &Submarine{Position: &Point{}}
	navigateSubmarine(s, commands)

	want := &Submarine{
		Aim:      10,
		Position: &Point{X: 15, Y: 60},
	}

	if diff := cmp.Diff(want, s); diff != "" {
		t.Errorf("readPoints() mismatch (-want +got):\n%s", diff)
	}
}
