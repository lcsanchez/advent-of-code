package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func readPoints(r io.Reader) ([]*Point, error) {
	points := make([]*Point, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()

		parts := strings.Split(s, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid instruction '%s'", s)
		}

		direction := parts[0]
		magnitude, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid magnitude '%s'", parts[1])
		}

		switch direction {
		case "forward":
			points = append(points, &Point{X: magnitude})
		case "down":
			points = append(points, &Point{Y: magnitude})
		case "up":
			points = append(points, &Point{Y: -magnitude})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points, nil
}

func navigate(points []*Point) *Point {
	position := &Point{
		X: 0,
		Y: 0,
	}
	for _, point := range points {
		position.Add(point)
	}

	return position
}
