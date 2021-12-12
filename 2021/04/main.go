package main

import (
	_ "embed"
	"io"
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
	return nil, nil, nil
}
