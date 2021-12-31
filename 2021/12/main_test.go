package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/test-input.txt
var testInput []byte

var expectedPaths = []*Path{
	{a: "start", b: "A"},
	{a: "start", b: "b"},
	{a: "A", b: "c"},
	{a: "A", b: "b"},
	{a: "b", b: "d"},
	{a: "A", b: "end"},
	{a: "b", b: "end"},
}

func TestReadInput(t *testing.T) {
	paths, err := readInput(bytes.NewReader(testInput))
	require.NoError(t, err)
	require.Equal(t, expectedPaths, paths)
}

func TestCreateCave(t *testing.T) {
	cave := createCave(expectedPaths)

	startNode := &Node{Label: "start"}
	aNode := &Node{Label: "A"}
	bNode := &Node{Label: "b"}
	cNode := &Node{Label: "c"}
	dNode := &Node{Label: "d"}
	endNode := &Node{Label: "end"}

	startNode.Nodes = []*Node{aNode, bNode}
	aNode.Nodes = []*Node{startNode, cNode, bNode, endNode}
	bNode.Nodes = []*Node{startNode, aNode, dNode, endNode}
	cNode.Nodes = []*Node{aNode}
	dNode.Nodes = []*Node{bNode}
	endNode.Nodes = []*Node{aNode, bNode}

	require.Equal(t, startNode, cave)
}

func TestFindPaths(t *testing.T) {
	start := createCave(expectedPaths)
	paths := make([][]*Node, 0)
	findPath(start, []*Node{start}, map[string]int{START: 1}, &paths, false)
	require.Equal(t, 10, len(paths))
}

func TestFindPaths_AllowDuplicateVisit(t *testing.T) {
	start := createCave(expectedPaths)
	paths := make([][]*Node, 0)
	findPath(start, []*Node{start}, map[string]int{START: 1}, &paths, true)
	require.Equal(t, 36, len(paths))
}
