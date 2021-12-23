package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
)

//go:embed testdata/input.txt
var input []byte

type Cave [][]int

func (c Cave) Get(p Point) (int, bool) {
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
	TOP             = Point{Y: -1}
	BOTTOM          = Point{Y: 1}
	LEFT            = Point{X: -1}
	RIGHT           = Point{X: 1}
	ADJACENT_POINTS = []Point{TOP, BOTTOM, LEFT, RIGHT}
)

func main() {
	cave, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findRiskLevelSum(cave))
	fmt.Println(findThreeLargestBasinMultiple(cave))
}

func readInput(r io.Reader) (Cave, error) {
	scanner := bufio.NewScanner(r)
	cave := [][]int{}

	for scanner.Scan() {
		s := scanner.Text()
		row := make([]int, len(s))

		for i, n := range s {
			var err error
			if row[i], err = strconv.Atoi(string(n)); err != nil {
				return nil, err
			}
		}

		cave = append(cave, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cave, nil
}

func findRiskLevelSum(cave Cave) int {
	sum := 0
	for rowIdx, row := range cave {
		for colIdx, num := range row {
			if isLowPoint(cave, Point{X: colIdx, Y: rowIdx}) {
				sum += num + 1
			}
		}
	}

	return sum
}

func isLowPoint(cave Cave, position Point) bool {
	currentValue, _ := cave.Get(position)

	for _, p := range ADJACENT_POINTS {
		value, ok := cave.Get(position.Add(p))
		if !ok {
			continue
		}

		if value <= currentValue {
			return false
		}
	}

	return true
}

func findThreeLargestBasinMultiple(cave Cave) int {
	basinSizes := []int{}
	for rowIdx, row := range cave {
		for colIdx, _ := range row {
			current := Point{X: colIdx, Y: rowIdx}
			if isLowPoint(cave, current) {
				basin := map[Point]bool{current: true}
				findBasinSize(cave, current, basin)

				basinSizes = append(basinSizes, len(basin))
			}
		}
	}

	sort.Ints(basinSizes)
	basinSizes = basinSizes[len(basinSizes)-3:]

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func findBasinSize(cave Cave, position Point, basin map[Point]bool) {
	currentValue, _ := cave.Get(position)

	for _, p := range ADJACENT_POINTS {
		posToCheck := position.Add(p)
		if _, ok := basin[posToCheck]; ok {
			continue
		}

		value, ok := cave.Get(posToCheck)
		if !ok || value == 9 {
			continue
		}

		if value > currentValue {
			basin[posToCheck] = true
			findBasinSize(cave, posToCheck, basin)
		}
	}
}
