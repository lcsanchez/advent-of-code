package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

func main() {
	moves, boards, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	score := playBingo(moves, boards)
	fmt.Println(score)

	losingScore := findLosingBoard(moves, boards)
	fmt.Println(losingScore)
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

func playBingo(moves []int, boards []*Board) int {
	for _, move := range moves {
		for _, board := range boards {
			board.MarkSpace(move)
			if board.isWinner() {
				return board.CalculateScore(move)
			}
		}
	}

	return 0
}

func findLosingBoard(moves []int, boards []*Board) int {
	winners := make(map[*Board]bool)

	for _, move := range moves {
		for _, board := range boards {
			if _, ok := winners[board]; ok {
				// board already won
				continue
			}

			board.MarkSpace(move)
			if board.isWinner() {
				winners[board] = true

				if len(winners) == len(boards) {
					return board.CalculateScore(move)
				}
			}
		}
	}

	return 0
}

type Space struct {
	Num    int
	Marked bool
}

type Board struct {
	Spaces [][]*Space
}

func (b *Board) MarkSpace(i int) {
	for _, row := range b.Spaces {
		for _, space := range row {
			if space.Num == i {
				space.Marked = true
			}
		}
	}
}

func (b *Board) isWinner() bool {
	for _, row := range b.Spaces {
		isWinner := true
		for _, space := range row {
			isWinner = isWinner && space.Marked
		}

		if isWinner {
			return true
		}
	}

	for col := 0; col < len(b.Spaces[0]); col++ {
		isWinner := true
		for row := 0; row < len(b.Spaces); row++ {
			isWinner = isWinner && b.Spaces[row][col].Marked
		}

		if isWinner {
			return true
		}
	}

	return false
}

func (b *Board) CalculateScore(i int) int {
	unmarkedSum := 0
	for _, row := range b.Spaces {
		for _, space := range row {
			if !space.Marked {
				unmarkedSum += space.Num
			}
		}
	}

	return unmarkedSum * i
}
