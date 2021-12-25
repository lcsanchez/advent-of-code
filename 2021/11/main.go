package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strconv"
)

//go:embed testdata/input.txt
var input []byte

type Grid [][]int

func (c Grid) Get(p Point) (int, bool) {
	if len(c) == 0 {
		return 0, false
	}

	isOutOfBounds := p.X < 0 || p.X >= len(c[0]) || p.Y < 0 || p.Y >= len(c)
	if isOutOfBounds {
		return 0, false
	}

	return c[p.Y][p.X], true
}

type Point struct {
	X, Y int
}

func (p Point) Add(o Point) Point {
	return Point{X: p.X + o.X, Y: p.Y + o.Y}
}

var (
	TOP_LEFT        = Point{X: -1, Y: -1}
	TOP             = Point{Y: -1}
	TOP_RIGHT       = Point{X: 1, Y: -1}
	LEFT            = Point{X: -1}
	RIGHT           = Point{X: 1}
	BOTTOM_LEFT     = Point{X: -1, Y: 1}
	BOTTOM          = Point{Y: 1}
	BOTTOM_RIGHT    = Point{X: 1, Y: 1}
	ADJACENT_POINTS = []Point{TOP_LEFT, TOP, TOP_RIGHT, LEFT, RIGHT, BOTTOM_LEFT, BOTTOM, BOTTOM_RIGHT}
)

const FLASH_POINT = 9

func main() {
	grid, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(flashesAfterSteps(copyGrid(grid), 100))
	fmt.Println(calculateSyncStep(copyGrid(grid)))
}

func readInput(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	grid := [][]int{}

	for scanner.Scan() {
		s := scanner.Text()
		row := make([]int, len(s))

		for i, n := range s {
			var err error
			if row[i], err = strconv.Atoi(string(n)); err != nil {
				return nil, err
			}
		}

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func flashesAfterSteps(grid [][]int, n int) int {
	flashes := 0
	for i := 0; i < n; i++ {
		flashes += step(grid)
	}

	return flashes
}

func calculateSyncStep(grid [][]int) int {
	var i int
	for i = 1; true; i++ {
		flashes := step(grid)

		if flashes == len(grid)*len(grid[0]) {
			break
		}
	}

	return i
}

func step(grid [][]int) int {
	flashes := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			current := Point{X: j, Y: i}
			if increaseEnergy(grid, current) {
				flashes += flash(grid, current)
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > FLASH_POINT {
				grid[i][j] = 0
			}
		}
	}

	return flashes
}

func increaseEnergy(grid [][]int, p Point) bool {
	if grid[p.Y][p.X] > FLASH_POINT {
		// already flashed
		return false
	}

	grid[p.Y][p.X]++

	return grid[p.Y][p.X] > FLASH_POINT
}

func flash(grid [][]int, current Point) int {
	flashes := 1
	for _, p := range ADJACENT_POINTS {
		p = current.Add(p)
		_, ok := Grid(grid).Get(p)
		if !ok {
			continue
		}

		if increaseEnergy(grid, p) {
			flashes += flash(grid, p)
		}
	}

	return flashes
}

func copyGrid(grid [][]int) [][]int {
	dupe := make([][]int, len(grid))
	for i := range grid {
		dupe[i] = make([]int, len(grid[i]))
		copy(dupe[i], grid[i])
	}

	return dupe
}
