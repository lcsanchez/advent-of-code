package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
)

//go:embed testdata/input.txt
var input []byte

type Closer struct {
	Points int
	Opener rune
}

var CLOSING_RUNES = map[rune]Closer{
	')': {Points: 3, Opener: '('},
	']': {Points: 57, Opener: '['},
	'}': {Points: 1197, Opener: '{'},
	'>': {Points: 25137, Opener: '<'},
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateTotalPoints(input))

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
			if s.Peek() != closer.Opener {
				return closer.Points
			}

			s.Pop()
			continue
		}

		s.Push(r)
	}

	return 0
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
