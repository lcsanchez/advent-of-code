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

type Tuple struct {
	First  int
	Second int
}

type Input struct {
	X *Tuple
	Y *Tuple
}

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(r io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	s := scanner.Text()
	parts := strings.Split(s[13:], ", ")
	fmt.Println(parts)

	xTuple, err := splitTuple(parts[0])
	if err != nil {
		return nil, err
	}

	yTuple, err := splitTuple(parts[1])
	if err != nil {
		return nil, err
	}

	return &Input{
		X: xTuple,
		Y: yTuple,
	}, nil
}

func splitTuple(s string) (*Tuple, error) {
	nums := strings.Split(s[2:], "..")

	first, err := strconv.Atoi(nums[0])
	if err != nil {
		return nil, err
	}

	second, err := strconv.Atoi(nums[1])
	if err != nil {
		return nil, err
	}

	if first > second {
		return &Tuple{
			First:  second,
			Second: first,
		}, nil
	}

	return &Tuple{
		First:  first,
		Second: second,
	}, nil
}
