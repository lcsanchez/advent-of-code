package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Space struct {
	Num    int
	Marked bool
}

type Board struct {
	Spaces [][]*Space
}

//go:embed testdata/input.txt
var input []byte

func main() {

}

func readInput(r io.Reader) ([]int, []*Board, error) {
	moves := make([]int, 0)
	boards := make([]*Board, 0)

	scanner := bufio.NewScanner(r)
	// Read moves
	if scanner.Scan() {
		s := scanner.Text()
		nums := strings.Split(s, ",")

		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}

			moves = append(moves, i)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	// Read boards
	var board *Board
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			if board != nil {
				boards = append(boards, board)
				board = nil
			}

			continue
		}

		if board == nil {
			board = &Board{}
		}

		row, err := parseRow(s)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing rows: %w", err)
		}

		board.Spaces = append(board.Spaces, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	boards = append(boards, board)

	return moves, boards, nil
}

func parseRow(s string) ([]*Space, error) {
	row := make([]*Space, 0)
	nums := strings.Fields(s)

	for _, num := range nums {

		i, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}

		row = append(row, &Space{Num: i})
	}

	return row, nil
}
