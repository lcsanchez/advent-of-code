package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
	"strconv"
)

//go:embed testdata/input.txt
var input []byte

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(r io.Reader) ([][]int, error) {
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
