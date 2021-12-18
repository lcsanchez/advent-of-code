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

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fish := input[0]
	calculateFishAfterDays(fish, 256)
	fmt.Println(countTotalFish(fish))
}

func readInput(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	fishDays := [][]int{}

	for scanner.Scan() {
		s := scanner.Text()
		fishDay := make([]int, 9)
		nums := strings.Split(s, ",")
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("parsing fish '%s': %w", s, err)
			}

			fishDay[num]++
		}

		fishDays = append(fishDays, fishDay)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fishDays, nil
}

func calculateFishAfterDays(fish []int, n int) {
	for i := 0; i < n; i++ {
		calculateFishAfterOneDay(fish)
	}
}

func calculateFishAfterOneDay(fish []int) {
	f := fish[0]
	fish[0] = 0

	for i := 1; i < len(fish); i++ {
		fish[i-1] = fish[i]
		fish[i] = 0
	}

	fish[6] += f
	fish[8] += f
}

func countTotalFish(fish []int) int {
	var total int
	for _, num := range fish {
		total += num
	}

	return total
}
