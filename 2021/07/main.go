package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

type FuelCostFunc func(float64) float64

func identity(n float64) float64 {
	return n
}

func sumToN(n float64) float64 {
	return (n * (n + 1)) / 2
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateOptimalHorizontalPostion(input, identity))
	fmt.Println(calculateOptimalHorizontalPostion(input, sumToN))
}

func readInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	positions := []int{}

	for scanner.Scan() {
		s := scanner.Text()
		nums := strings.Split(s, ",")
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("parsing position '%s': %w", s, err)
			}

			positions = append(positions, num)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sort.Ints(positions)

	return positions, nil
}

func calculateOptimalHorizontalPostion(positions []int, fuelCostFunc FuelCostFunc) (int, int) {
	optimalPosition := 0
	optimalFuelCost := calculateFuelCost(positions, 0, fuelCostFunc)

	for i := 1; i < positions[len(positions)-1]; i++ {
		currentFuelCost := calculateFuelCost(positions, i, fuelCostFunc)
		if optimalFuelCost > currentFuelCost {
			optimalFuelCost = currentFuelCost
			optimalPosition = i
		}
	}

	return optimalPosition, int(optimalFuelCost)
}

func calculateFuelCost(positions []int, position int, fuelCostFunc FuelCostFunc) float64 {
	var fuelCost float64

	for j := 0; j < len(positions); j++ {
		fuelCost += fuelCostFunc(math.Abs(float64(position - positions[j])))
	}

	return fuelCost
}
