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

type FoldAxis int

const (
	FoldAxisX FoldAxis = iota
	FoldAxisY
)

type Fold struct {
	FoldAxis FoldAxis
	Value    int
}

type Input struct {
	Points []Point
	Folds  []Fold
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	points := input.Points
	points = fold(points, input.Folds[0])
	fmt.Println(len(points))

	// Part 2
	points = input.Points
	for _, f := range input.Folds {
		points = fold(points, f)
	}
	fmt.Println(len(points))
	print(points)
}

func readInput(r io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(r)
	input := &Input{
		Points: []Point{},
		Folds:  []Fold{},
	}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}

		point, err := readPoint(s)
		if err != nil {
			return nil, err
		}

		input.Points = append(input.Points, point)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		s := scanner.Text()
		fold, err := readFold(s)
		if err != nil {
			return nil, err
		}

		input.Folds = append(input.Folds, fold)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func readPoint(s string) (Point, error) {
	nums := strings.Split(s, ",")

	x, err := strconv.Atoi(nums[0])
	if err != nil {
		return Point{}, fmt.Errorf("parsing x value from point: %s", s)
	}

	y, err := strconv.Atoi(nums[1])
	if err != nil {
		return Point{}, fmt.Errorf("parsing y value from point: %s", s)
	}

	return Point{X: x, Y: y}, nil
}

func readFold(s string) (Fold, error) {
	strs := strings.Fields(s)
	fold := strings.Split(strs[2], "=")

	var axis FoldAxis
	switch fold[0] {
	case "x":
		axis = FoldAxisX
	case "y":
		axis = FoldAxisY
	default:
		return Fold{}, fmt.Errorf("parsing fold axis: %s", fold[0])
	}

	value, err := strconv.Atoi(fold[1])
	if err != nil {
		return Fold{}, fmt.Errorf("parsing fold value: %s", fold[1])
	}

	return Fold{FoldAxis: axis, Value: value}, nil
}

func fold(points []Point, fold Fold) []Point {
	pointMap := map[Point]bool{}

	for _, point := range points {
		switch fold.FoldAxis {
		case FoldAxisX:
			if point.X == fold.Value {
				continue
			}

			newPoint := Point{X: point.X, Y: point.Y}
			if point.X > fold.Value {
				newPoint.X = fold.Value - (point.X - fold.Value)
			}

			pointMap[newPoint] = true
		case FoldAxisY:
			if point.Y == fold.Value {
				continue
			}

			newPoint := Point{X: point.X, Y: point.Y}
			if point.Y > fold.Value {
				newPoint.Y = fold.Value - (point.Y - fold.Value)
			}

			pointMap[newPoint] = true
		}
	}

	newPoints := make([]Point, 0, len(pointMap))
	for k, _ := range pointMap {
		newPoints = append(newPoints, k)
	}

	return newPoints
}

func findDimensions(points []Point) (int, int) {
	maxX, maxY := 0, 0

	for _, p := range points {
		maxX = int(math.Max(float64(p.X), float64(maxX)))
		maxY = int(math.Max(float64(p.Y), float64(maxY)))
	}

	return maxX, maxY
}

func print(points []Point) {
	x, y := findDimensions(points)
	pointMap := map[Point]bool{}
	for _, p := range points {
		pointMap[p] = true
	}

	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			r := "."
			if _, ok := pointMap[Point{X: j, Y: i}]; ok {
				r = "#"
			}

			fmt.Print(r)
		}
		fmt.Println()
	}
}
