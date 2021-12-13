package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

func main() {

}

func readInput(r io.Reader) ([]*Line, error) {
	scanner := bufio.NewScanner(r)
	lines := make([]*Line, 0)

	for scanner.Scan() {
		s := scanner.Text()
		points := strings.Split(s, " -> ")

		start, err := readPoint(points[0])
		if err != nil {
			return nil, err
		}
		end, err := readPoint(points[1])
		if err != nil {
			return nil, err
		}

		lines = append(lines, &Line{Start: start, End: end})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func readPoint(s string) (*Point, error) {
	nums := strings.Split(s, ",")
	x, err := strconv.Atoi(nums[0])
	if err != nil {
		return nil, fmt.Errorf("reading point x value '%s': %w", nums[0], err)
	}

	y, err := strconv.Atoi(nums[1])
	if err != nil {
		return nil, fmt.Errorf("reading point x value '%s': %w", nums[1], err)
	}

	return &Point{X: x, Y: y}, nil
}

type Point struct {
	X, Y int
}

type Line struct {
	Start *Point
	End   *Point
}
