package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		report = append(report, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	powerConsumption := calculatePowerConsumption(report)
	fmt.Println(powerConsumption)

	lifeSupportRating, err := calculateLifeSupportRating(report)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lifeSupportRating)
}

func calculatePowerConsumption(report []string) int {
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

	return gamma * epsilon
}

func calculateLifeSupportRating(report []string) (int, error) {
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
