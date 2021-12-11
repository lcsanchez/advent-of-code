package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
)

//go:embed testdata/input.txt
var input []byte

func main() {
	// Part 01
	powerConsumption, err := calculatePowerConsumption(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(powerConsumption)

	// Part 02
	lifeSupportRating, err := calculateLifeSupportRating(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lifeSupportRating)
}

func readReport(r io.Reader) ([]string, error) {
	report := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		report = append(report, s)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return report, nil
}

func calculatePowerConsumption(r io.Reader) (int, error) {
	report, err := readReport(r)
	if err != nil {
		return 0, err
	}

	bitCount := len(report[0])
	bitTracker := make([]int, bitCount)

	for _, reading := range report {
		for j, c := range reading {
			if c == '1' {
				bitTracker[j]++
			}
		}
	}

	gamma, epsilon := 0, 0
	for i, count := range bitTracker {
		if count > (len(report) / 2) {
			gamma |= (1 << (bitCount - i - 1))
		} else {
			epsilon |= (1 << (bitCount - i - 1))
		}
	}

	return gamma * epsilon, nil
}

type RatingFilterFunc func(report []string, pivot int) []string

func O2RatingFilterFunc(report []string, pivot int) []string {
	if (len(report) - pivot) < pivot {
		return report[:pivot]
	}

	return report[pivot:]
}

func CO2RatingFilterFunc(report []string, pivot int) []string {
	if (len(report) - pivot) < pivot {
		return report[pivot:]
	}

	return report[:pivot]
}

func calculateLifeSupportRating(r io.Reader) (int, error) {
	report, err := readReport(r)
	if err != nil {
		return 0, err
	}

	sort.Strings(report)

	oxygenRating, err := calculateRating(report, 0, O2RatingFilterFunc)
	if err != nil {
		return 0, fmt.Errorf("calculating oxygen rating: %v", err)
	}
	co2Rating, err := calculateRating(report, 0, CO2RatingFilterFunc)
	if err != nil {
		return 0, fmt.Errorf("calculating co2 scubber rating: %v", err)
	}

	return oxygenRating * co2Rating, nil

}

func calculateRating(report []string, bitIndex int, f RatingFilterFunc) (int, error) {
	if len(report) == 1 {
		rating, err := strconv.ParseInt(report[0], 2, 32)
		if err != nil {
			return 0, err
		}

		return int(rating), nil
	}

	pivot := findPivot(report, bitIndex)

	return calculateRating(f(report, pivot), bitIndex+1, f)
}

func findPivot(report []string, bitIndex int) int {
	for i, s := range report {
		if s[bitIndex] == '1' {
			return i
		}
	}

	return len(report)
}
