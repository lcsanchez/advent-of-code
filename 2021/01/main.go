package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := countDepthIncrease(nums)
	log.Println(count)

	count = countDepthGroupIncrease(nums, 3)
	log.Println(count)
}

func countDepthIncrease(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}

	return count
}

func countDepthGroupIncrease(nums []int, groupSize int) int {
	return countDepthIncrease(makeGroups(nums, groupSize))
}

func makeGroups(nums []int, groupSize int) []int {
	if len(nums) < groupSize {
		return []int{}
	}

	groups := make([]int, len(nums)-(groupSize-1))
	for i := 0; i < len(groups); i++ {
		for j := 0; j < groupSize; j++ {
			groups[i] += nums[i+j]
		}
	}

	return groups
}
