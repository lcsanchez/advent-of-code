package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

type Point struct {
	X, Y int
}

func (p *Point) Equal(o *Point) bool {
	return p.X == o.X && p.Y == o.Y
}

func (p *Point) Add(o *Point) {
	p.X += o.X
	p.Y += o.Y
}

type Line struct {
	Start *Point
	End   *Point
}

func main() {
	lines, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	maxPoint := calculateBoardMaxPoint(lines)
	grid := initializeGrid(maxPoint.X, maxPoint.Y)
	drawLines(grid, lines, true)

	fmt.Println(findOverlappingCount(grid))

	grid2 := initializeGrid(maxPoint.X, maxPoint.Y)
	drawLines(grid2, lines, false)
	fmt.Println(findOverlappingCount(grid2))
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

func calculateBoardMaxPoint(lines []*Line) *Point {
	maxX, maxY := 0, 0
	for _, line := range lines {
		maxX = int(math.Max(float64(maxX), math.Max(float64(line.Start.X), float64(line.End.X))))
		maxY = int(math.Max(float64(maxY), math.Max(float64(line.Start.Y), float64(line.End.Y))))
	}

	return &Point{X: maxX, Y: maxY}
}

func drawLines(grid [][]int, lines []*Line, ignoreDiags bool) {
	for _, line := range lines {
		if ignoreDiags && line.Start.X != line.End.X && line.Start.Y != line.End.Y {
			continue
		}

		inc := &Point{}
		if line.Start.X < line.End.X {
			inc.X = 1
		} else if line.Start.X > line.End.X {
			inc.X = -1
		}

		if line.Start.Y < line.End.Y {
			inc.Y = 1
		} else if line.Start.Y > line.End.Y {
			inc.Y = -1
		}

		drawLine(grid, line, inc)
	}
}

func drawLine(grid [][]int, line *Line, inc *Point) {
	current := &Point{X: line.Start.X, Y: line.Start.Y}

	for !current.Equal(line.End) {
		grid[current.Y][current.X] += 1

		current.Add(inc)
	}

	grid[current.Y][current.X] += 1
}

func initializeGrid(x, y int) [][]int {
	grid := make([][]int, y+1)
	for i := range grid {
		grid[i] = make([]int, x+1)
	}

	return grid
}

func findOverlappingCount(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, num := range row {
			if num > 1 {
				count++
			}
		}
	}

	return count
}
