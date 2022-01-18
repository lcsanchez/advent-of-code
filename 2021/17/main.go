package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

type Tuple struct {
	First  int
	Second int
}

type Input struct {
	X *Tuple
	Y *Tuple
}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateSummation(input.Y.First - 1))
	fmt.Println(countStartingVelocities(input))
}

func readInput(r io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	s := scanner.Text()
	parts := strings.Split(s[13:], ", ")
	fmt.Println(parts)

	xTuple, err := splitTuple(parts[0])
	if err != nil {
		return nil, err
	}

	yTuple, err := splitTuple(parts[1])
	if err != nil {
		return nil, err
	}

	return &Input{
		X: xTuple,
		Y: yTuple,
	}, nil
}

func splitTuple(s string) (*Tuple, error) {
	nums := strings.Split(s[2:], "..")

	first, err := strconv.Atoi(nums[0])
	if err != nil {
		return nil, err
	}

	second, err := strconv.Atoi(nums[1])
	if err != nil {
		return nil, err
	}

	return &Tuple{
		First:  first,
		Second: second,
	}, nil
}

func findSummationInRange(t *Tuple) (int, int) {
	sum := 0
	idx := 1
	for {
		sum += idx

		if sum > t.First {
			break
		}

		idx++
	}

	return idx, sum
}

func calculateSummation(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}

func countStartingVelocities(input *Input) int {
	velocities := (int(math.Abs(float64(input.X.Second-input.X.First))) + 1) * (int(math.Abs(float64(input.Y.Second-input.Y.First))) + 1)

	xCurrent := (input.X.Second / 2) + (input.X.Second % 2)
	for i := xCurrent; i > 0; i-- {
		sum := i
		steps := 1
		stepMin := -1
		var stepMax int
		for {
			if input.X.Second < sum {
				stepMax = steps - 1
				break
			}

			if stepMin == -1 && input.X.First <= sum {
				stepMin = steps
			}

			if i == steps {
				if input.X.First < sum && sum < input.X.Second {
					stepMax = math.MaxInt64
				}

				break
			}

			sum += i - steps
			steps++

		}

		if stepMin != -1 {
			velocities += calculateYVelocitiesInSteps(input.Y, stepMin, stepMax)
		}

	}

	return velocities
}

// These could be precalculated instead of for each different x
func calculateYVelocitiesInSteps(yRange *Tuple, stepMin int, stepMax int) int {
	yVelocities := 0
	for i := int(math.Abs(float64(yRange.First))) - 1; i >= yRange.First; i-- {
		currentY := i
		for step := 1; step <= stepMax; step++ {
			if currentY < yRange.First {
				break
			}

			if step >= stepMin && currentY >= yRange.First && currentY <= yRange.Second {
				yVelocities++
				break
			}

			currentY += i - step
		}
	}

	return yVelocities
}
