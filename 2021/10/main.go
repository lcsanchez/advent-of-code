package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"sort"
)

//go:embed testdata/input.txt
var input []byte

type Balancer struct {
	Points  int
	Balance rune
}

var CLOSING_RUNES = map[rune]Balancer{
	')': {Points: 3, Balance: '('},
	']': {Points: 57, Balance: '['},
	'}': {Points: 1197, Balance: '{'},
	'>': {Points: 25137, Balance: '<'},
}

var OPENING_RUNES = map[rune]Balancer{
	'(': {Points: 1, Balance: ')'},
	'[': {Points: 2, Balance: ']'},
	'{': {Points: 3, Balance: '}'},
	'<': {Points: 4, Balance: '>'},
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateTotalPoints(input))
	fmt.Println(calculateMiddleMissingCloserScore(input))
}

func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func calculateTotalPoints(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += calculatePoints(line)
	}

	return sum
}

func calculatePoints(line string) int {
	s := NewStack()
	for _, r := range line {
		if closer, ok := CLOSING_RUNES[r]; ok {
			if s.Peek() != closer.Balance {
				return closer.Points
			}

			s.Pop()
			continue
		}

		s.Push(r)
	}

	return 0
}

func calculateMiddleMissingCloserScore(lines []string) int {
	scores := []int{}

	for _, line := range lines {
		s, ok := calculateMissingClosers(line)
		if !ok {
			continue
		}

		scores = append(scores, calculateMissingCloserScore(s))
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func calculateMissingClosers(line string) (*Stack, bool) {
	s := NewStack()
	for _, r := range line {
		if closer, ok := CLOSING_RUNES[r]; ok {
			if s.Peek() != closer.Balance {
				return nil, false
			}

			s.Pop()
			continue
		}

		s.Push(r)
	}

	return s, true
}

func calculateMissingCloserScore(s *Stack) int {
	score := 0
	for {
		r, ok := s.Pop()
		if !ok {
			break
		}

		score = score*5 + OPENING_RUNES[r].Points
	}

	return score
}

func NewStack() *Stack {
	return &Stack{stack: []rune{}}
}

type Stack struct {
	stack []rune
}

func (s *Stack) Push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *Stack) Peek() rune {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}

	r := s.Peek()
	s.stack = s.stack[:len(s.stack)-1]

	return r, true
}
