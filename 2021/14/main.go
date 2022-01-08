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
	Insert rune
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	min, max := applyPolymer(input.Template, input.Rules, 40)

	fmt.Println(max - min)
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
		Insert: []rune(parts[1])[0],
	}, nil
}

func applyPolymer(template string, rules []*Rule, times int) (int, int) {
	ruleMap := map[string]rune{}

	for _, rule := range rules {
		ruleMap[rule.Match] = rule.Insert
	}

	pairs := createPairs(template)

	for i := 0; i < times; i++ {
		newPairs := map[string]int{}
		for pair, count := range pairs {
			if insert, ok := ruleMap[pair]; ok {
				runePair := []rune(pair)

				newFirstPair := string([]rune{runePair[0], insert})
				newSecondPair := string([]rune{insert, runePair[1]})

				newPairs[newFirstPair] += count
				newPairs[newSecondPair] += count
			} else {
				newPairs[pair] += pairs[pair]
			}
		}

		pairs = newPairs
	}

	runeCount := pairsToRuneCount(pairs)

	// we need to also account for the last char in the template
	runeCount[[]rune(template)[len(template)-1]]++

	return minMaxRuneCount(runeCount)
}

func createPairs(s string) map[string]int {
	pairs := map[string]int{}

	var previous rune
	for _, r := range s {
		if previous != 0 {
			pairs[string([]rune{previous, r})]++
		}
		previous = r
	}

	return pairs
}

func pairsToRuneCount(pairs map[string]int) map[rune]int {
	runeCount := map[rune]int{}

	for pair, count := range pairs {
		runes := []rune(pair)

		runeCount[runes[0]] += count
	}

	return runeCount
}

func minMaxRuneCount(runes map[rune]int) (int, int) {
	var max, min int
	for _, value := range runes {
		if min == 0 || min > value {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return min, max
}
