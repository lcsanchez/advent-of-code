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
	moves, boards, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)

	wantMoves := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	require.Equal(t, wantMoves, moves)

	wantBoards := []Board{
		{
			Spaces: [][]*Space{
				{{Num: 22}, {Num: 13}, {Num: 17}, {Num: 11}, {Num: 0}},
				{{Num: 8}, {Num: 2}, {Num: 23}, {Num: 4}, {Num: 24}},
				{{Num: 21}, {Num: 9}, {Num: 14}, {Num: 16}, {Num: 7}},
				{{Num: 6}, {Num: 10}, {Num: 3}, {Num: 18}, {Num: 5}},
				{{Num: 1}, {Num: 12}, {Num: 20}, {Num: 15}, {Num: 19}},
			},
		},
		{
			Spaces: [][]*Space{
				{{Num: 3}, {Num: 15}, {Num: 0}, {Num: 2}, {Num: 22}},
				{{Num: 9}, {Num: 18}, {Num: 13}, {Num: 17}, {Num: 5}},
				{{Num: 19}, {Num: 8}, {Num: 7}, {Num: 25}, {Num: 23}},
				{{Num: 20}, {Num: 11}, {Num: 10}, {Num: 24}, {Num: 4}},
				{{Num: 14}, {Num: 21}, {Num: 16}, {Num: 12}, {Num: 6}},
			},
		},
		{
			Spaces: [][]*Space{
				{{Num: 14}, {Num: 21}, {Num: 17}, {Num: 24}, {Num: 4}},
				{{Num: 10}, {Num: 16}, {Num: 15}, {Num: 9}, {Num: 19}},
				{{Num: 18}, {Num: 8}, {Num: 23}, {Num: 26}, {Num: 20}},
				{{Num: 22}, {Num: 11}, {Num: 13}, {Num: 6}, {Num: 5}},
				{{Num: 2}, {Num: 0}, {Num: 12}, {Num: 3}, {Num: 7}},
			},
		},
	}
	require.Equal(t, wantBoards, boards)
}
