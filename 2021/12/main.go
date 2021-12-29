package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode"
)

//go:embed testdata/input.txt
var input []byte

const (
	START = "start"
	END   = "end"
)

type Path struct {
	a string
	b string
}

type Node struct {
	Label string
	Nodes map[string]*Node
}

func NewNode(label string) *Node {
	return &Node{
		Label: label,
		Nodes: make(map[string]*Node),
	}
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	start := createCave(input)
	paths := make([][]*Node, 0)
	findPath(start, []*Node{start}, map[string]int{START: 1}, &paths, true)
	fmt.Println(len(paths))
}

func readInput(r io.Reader) ([]*Path, error) {
	scanner := bufio.NewScanner(r)
	paths := []*Path{}

	for scanner.Scan() {
		s := scanner.Text()
		nodes := strings.Split(s, "-")

		paths = append(paths, &Path{
			a: nodes[0],
			b: nodes[1],
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return paths, nil
}

func createCave(paths []*Path) *Node {
	start := NewNode(START)
	end := NewNode(END)

	nodesByLabel := map[string]*Node{
		START: start,
		END:   end,
	}

	for _, path := range paths {
		aNode, ok := nodesByLabel[path.a]
		if !ok {
			aNode = NewNode(path.a)
			nodesByLabel[path.a] = aNode
		}

		bNode, ok := nodesByLabel[path.b]
		if !ok {
			bNode = NewNode(path.b)
			nodesByLabel[path.b] = bNode
		}

		aNode.Nodes[bNode.Label] = bNode
		bNode.Nodes[aNode.Label] = aNode
	}

	return start
}

func findPath(current *Node, visited []*Node, visitedMap map[string]int, paths *[][]*Node, allowMultipleVisits bool) {
	if current.Label == END {
		path := make([]*Node, len(visited))
		copy(path, visited)
		*paths = append(*paths, path)
		return
	}

	for _, next := range current.Nodes {
		var isSecondVisit bool
		if _, ok := visitedMap[next.Label]; ok {
			if !allowMultipleVisits || next.Label == START {
				continue
			}
			isSecondVisit = true
			allowMultipleVisits = false
		}

		visited = append(visited, next)

		if isSmallCave(next.Label) {
			visitedMap[next.Label]++
		}

		findPath(next, visited, visitedMap, paths, allowMultipleVisits)

		visited = visited[:len(visited)-1]
		if i, ok := visitedMap[next.Label]; ok {
			if i <= 1 {
				delete(visitedMap, next.Label)
			} else {
				visitedMap[next.Label]--
			}
		}

		if isSecondVisit {
			allowMultipleVisits = true
		}
	}
}

func isSmallCave(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
