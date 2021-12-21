package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
	"strings"
)

type Entry struct {
	SignalPatterns []string
	OutputValues   []string
}

//go:embed testdata/input.txt
var input []byte

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

}

func readInput(r io.Reader) ([]*Entry, error) {
	scanner := bufio.NewScanner(r)
	entries := []*Entry{}

	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, "|")

		entries = append(entries, &Entry{
			SignalPatterns: strings.Fields(parts[0]),
			OutputValues:   strings.Fields(parts[1]),
		})

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
