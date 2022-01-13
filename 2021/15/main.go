package main

import (
	"bufio"
	"bytes"
	"container/heap"
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
	TOP             = Point{Y: -1}
	LEFT            = Point{X: -1}
	RIGHT           = Point{X: 1}
	BOTTOM          = Point{Y: 1}
	ADJACENT_POINTS = []Point{TOP, RIGHT, BOTTOM, LEFT}
)

type Path struct {
	Current Point
	Value   int
	index   int
}

type PathHeap []*Path

func (h PathHeap) Len() int {
	return len(h)
}
func (h PathHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}
func (h PathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *PathHeap) Push(x interface{}) {
	*h = append(*h, x.(*Path))
}

func (h *PathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	item.index = -1
	return item
}

func main() {
	cave, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findLowestRiskPath(cave))

	largeCave := buildLargeCave(cave, 5)
	fmt.Println(findLowestRiskPath(largeCave))
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

func buildLargeCave(cave [][]int, n int) [][]int {
	largeCave := [][]int{}

	for repeat := 0; repeat < n; repeat++ {
		colLen := len(cave)
		for row := 0; row < colLen; row++ {
			largeCave = append(largeCave, make([]int, len(cave[row])*n))
			for colRepeat := 0; colRepeat < n; colRepeat++ {
				rowLen := len(cave[row])
				for col := 0; col < rowLen; col++ {
					newValue := (cave[row][col] + repeat + colRepeat)
					largeCave[repeat*colLen+row][colRepeat*rowLen+col] = newValue - ((newValue / 10) * 9)
				}
			}
		}
	}

	return largeCave
}

func findLowestRiskPath(grid Grid) int {
	visited := map[Point]bool{}
	h := &PathHeap{}
	heap.Push(h, &Path{Current: Point{X: 0, Y: 0}, Value: 0})

	for h.Len() > 0 {
		path := heap.Pop(h).(*Path)

		if path.Current.X == len(grid[0])-1 && path.Current.Y == len(grid)-1 {
			return path.Value
		}

		for _, point := range ADJACENT_POINTS {
			next := path.Current.Add(point)

			if visited[next] {
				continue
			}

			value, ok := grid.Get(next)
			if !ok {
				continue
			}

			heap.Push(h, &Path{Current: next, Value: path.Value + value})
		}

		visited[path.Current] = true
	}

	return -1
}
