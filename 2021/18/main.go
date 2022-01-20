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

type Node struct {
	left  *Node
	right *Node
	data  int
}

func (n *Node) clone() *Node {
	new := &Node{
		data: n.data,
	}

	if n.left != nil {
		new.left = n.left.clone()
	}

	if n.right != nil {
		new.right = n.right.clone()
	}

	return new
}

func (n *Node) isData() bool {
	if n == nil {
		return false
	}

	return n.left == nil && n.right == nil
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	total := input[0]
	for i := 1; i < len(input); i++ {
		total = reduce(sum(total.clone(), input[i].clone()))
	}

	fmt.Println(magnitude(total))

	var m int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}

			current := magnitude(reduce(sum(input[i].clone(), input[j].clone())))
			if current > m {
				m = current
			}
		}
	}

	fmt.Println(m)
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

func magnitude(node *Node) int {
	if node.isData() {
		return node.data
	}

	return 3*magnitude(node.left) + 2*magnitude(node.right)
}

func print(node *Node) string {
	var sb strings.Builder
	printHelper(node, &sb)

	return sb.String()
}

func printHelper(node *Node, sb *strings.Builder) {
	if node.isData() {
		sb.WriteString(strconv.Itoa(node.data))
		return
	}

	sb.WriteRune('[')
	printHelper(node.left, sb)
	sb.WriteRune(',')
	printHelper(node.right, sb)
	sb.WriteRune(']')
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

func sum(left *Node, right *Node) *Node {
	return &Node{
		left:  left,
		right: right,
	}
}

func reduce(root *Node) *Node {
	for {
		for {
			exploder := findExploder(root, 0)
			if exploder == nil {
				break
			}

			explode(root, exploder)
		}

		splitter := findSplitter(root)
		if splitter == nil {
			break
		}
		split(splitter)
	}

	return root
}

func findSplitter(node *Node) *Node {
	if node.isData() {
		if node.data > 9 {
			return node
		}

		return nil
	}

	splitter := findSplitter(node.left)
	if splitter != nil {
		return splitter
	}

	splitter = findSplitter(node.right)
	if splitter != nil {
		return splitter
	}

	return nil
}

func split(node *Node) {
	node.left = &Node{data: node.data / 2}
	node.right = &Node{data: node.data/2 + (node.data % 2)}
	node.data = 0
}

func explode(root *Node, exploder *Node) {
	left := findLeft(root, exploder)
	if left != nil {
		left.data += exploder.left.data
	}

	right := findRight(root, exploder)
	if right != nil {
		right.data += exploder.right.data
	}

	exploder.left = nil
	exploder.right = nil
	exploder.data = 0
}

func findExploder(node *Node, depth int) *Node {
	if node.isData() {
		return nil
	}

	if node.left.isData() && node.right.isData() {
		if depth < 4 {
			return nil
		}

		return node
	}

	exploder := findExploder(node.left, depth+1)
	if exploder != nil {
		return exploder
	}
	exploder = findExploder(node.right, depth+1)

	return exploder
}

func findLeft(root *Node, target *Node) *Node {
	var result *Node
	findLeftHelper(root, target, &result)

	return result
}

func findLeftHelper(node *Node, target *Node, result **Node) bool {
	if node == target {
		return true
	}

	if node.isData() {
		*result = node
		return false
	}

	return findLeftHelper(node.left, target, result) ||
		findLeftHelper(node.right, target, result)
}

func findRight(root *Node, target *Node) *Node {
	var result *Node
	findRightHelper(root, target, &result)

	return result
}

func findRightHelper(node *Node, target *Node, result **Node) bool {
	if node == target {
		return true
	}

	if node.isData() {
		*result = node
		return false
	}

	return findRightHelper(node.right, target, result) ||
		findRightHelper(node.left, target, result)
}
