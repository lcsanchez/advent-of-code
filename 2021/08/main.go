package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

// https://stackoverflow.com/a/22698017
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

type Entry struct {
	SignalPatterns []string
	OutputValues   []string
}

//go:embed testdata/input.txt
var input []byte

func main() {
	entries, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	_, total := countIdentifiableValues(entries)
	fmt.Println(total)

	output, err := sumOutputValues(entries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
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

func countIdentifiableValues(entries []*Entry) (map[int]int, int) {
	counts := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}

	for _, entry := range entries {
		for _, value := range entry.OutputValues {
			switch len(value) {
			case 2:
				counts[1]++
			case 4:
				counts[4]++
			case 3:
				counts[7]++
			case 7:
				counts[8]++
			}
		}
	}

	return counts, counts[1] + counts[4] + counts[7] + counts[8]
}

func sumOutputValues(entries []*Entry) (int, error) {
	totalOutput := 0
	for _, entry := range entries {
		decodedPatterns := decodeSignalPatterns(entry)

		output, err := decodeOutputValues(decodedPatterns, entry.OutputValues)
		if err != nil {
			return 0, err
		}
		totalOutput += output
	}

	return totalOutput, nil
}

func decodeOutputValues(signals map[string]int, values []string) (int, error) {
	output := 0
	for _, o := range values {
		r := []rune(o)
		sort.Sort(sortRunes(r))

		num, ok := signals[string(r)]
		if !ok {
			fmt.Errorf("unable to decode output: %s", o)
		}
		output = output*10 + num
	}

	return output, nil
}

func decodeSignalPatterns(entry *Entry) map[string]int {
	signalPatternsByLen := sortedValuesByLen(entry)
	signalPatternOne := signalPatternsByLen[2][0]
	signalPatternFour := signalPatternsByLen[4][0]
	signalPatternSeven := signalPatternsByLen[3][0]
	signalPatternEight := signalPatternsByLen[7][0]
	solvedSignalPatterns := map[string]int{
		string(signalPatternOne):   1,
		string(signalPatternFour):  4,
		string(signalPatternSeven): 7,
		string(signalPatternEight): 8,
	}

	for _, value := range signalPatternsByLen[5] {
		unionWithSeven := union(signalPatternSeven, value)
		if len(unionWithSeven) == 3 {
			solvedSignalPatterns[string(value)] = 3
			continue
		}

		unionWithFour := union(signalPatternFour, value)
		if len(unionWithFour) == 2 {
			solvedSignalPatterns[string(value)] = 2
			continue
		}

		solvedSignalPatterns[string(value)] = 5
	}

	for _, value := range signalPatternsByLen[6] {
		unionWithFour := union(signalPatternFour, value)
		if len(unionWithFour) == 4 {
			solvedSignalPatterns[string(value)] = 9
			continue
		}

		unionWithSeven := union(signalPatternSeven, value)
		if len(unionWithSeven) == 3 {
			solvedSignalPatterns[string(value)] = 0
			continue
		}

		solvedSignalPatterns[string(value)] = 6
	}

	return solvedSignalPatterns
}

func sortedValuesByLen(entry *Entry) map[int][][]rune {
	valuesByLen := map[int][][]rune{}

	for _, value := range entry.SignalPatterns {
		r := []rune(value)
		sort.Sort(sortRunes(r))

		valuesByLen[len(value)] = append(valuesByLen[len(value)], r)
	}

	return valuesByLen
}

func union(a, b []rune) []rune {
	aSet := map[rune]bool{}

	for _, aVal := range a {
		aSet[aVal] = true
	}

	resultSet := []rune{}
	for _, bVal := range b {
		if _, ok := aSet[bVal]; ok {
			resultSet = append(resultSet, bVal)
		}
	}

	return resultSet
}
