package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
)

//go:embed testdata/input.txt
var input []byte

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}
