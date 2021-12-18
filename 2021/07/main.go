package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

const (
	AdultLifecycleDays    = 6
	OffspingLifecycleDays = 8
)

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

}

func readInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	positions := []int{}

	for scanner.Scan() {
		s := scanner.Text()
		nums := strings.Split(s, ",")
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("parsing position '%s': %w", s, err)
			}

			positions = append(positions, num)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func calculateOptimalHorizontalPostion(positions []int) int {
	return 0
}
