package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

type Input struct {
	Template string
	Rules    []*Rule
}

type Rule struct {
	Match  string
	Insert string
}

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(r io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(r)
	input := &Input{
		Rules: []*Rule{},
	}

	if scanner.Scan() {
		input.Template = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			continue
		}
		rule, err := readRule(s)
		if err != nil {
			return nil, err
		}

		input.Rules = append(input.Rules, rule)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func readRule(s string) (*Rule, error) {
	parts := strings.Split(s, " -> ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("error parsing rule")
	}

	return &Rule{
		Match:  parts[0],
		Insert: parts[1],
	}, nil
}
