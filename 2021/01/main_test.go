package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountDepthIncrease(t *testing.T) {
	count := countDepthIncrease([]int{
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
	})

	require.Equal(t, 7, count)
}

func TestMakeGroups(t *testing.T) {
	got := makeGroups([]int{
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
	}, 3)

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
	count := countDepthGroupIncrease([]int{
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
	}, 3)

	require.Equal(t, 5, count)
}
