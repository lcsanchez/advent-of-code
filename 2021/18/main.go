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

type Node struct {
	left  *Node
	right *Node
	data  int
}

func main() {
	_, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(r io.Reader) ([]*Node, error) {
	scanner := bufio.NewScanner(r)

	nodes := []*Node{}
	for scanner.Scan() {
		s := scanner.Text()
		node, err := readNode(s)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nodes, nil
}

func readNode(s string) (*Node, error) {
	if len(s) == 1 {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		return &Node{
			data: num,
		}, nil
	}

	s = s[1 : len(s)-1]
	leftStr, rightStr := splitOnFreeComma(s)

	left, err := readNode(leftStr)
	if err != nil {
		return nil, err
	}

	right, err := readNode(rightStr)
	if err != nil {
		return nil, err
	}

	return &Node{
		left:  left,
		right: right,
	}, nil
}

func splitOnFreeComma(s string) (string, string) {
	groupCount := 0
	for idx, r := range s {
		if r == ',' && groupCount == 0 {
			return s[0:idx], s[idx+1:]
		}

		if r == '[' {
			groupCount++
		}

		if r == ']' {
			groupCount--
		}
	}

	return "", ""
}
