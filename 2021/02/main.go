package main

import (
	"fmt"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Add(po *Point) {
	p.X += po.X
	p.Y += po.Y
}

func main() {
	answerOne, err := partOne()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answerOne)

	answerTwo, err := partTwo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answerTwo)
}

func partOne() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	points, err := readPoints(file)
	if err != nil {
		return 0, err
	}

	position := navigate(points)
	return position.X * position.Y, nil
}

func partTwo() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	commands, err := readCommands(file)
	if err != nil {
		return 0, err
	}

	submarine := &Submarine{
		Position: &Point{},
	}
	navigateSubmarine(submarine, commands)

	return submarine.Position.X * submarine.Position.Y, nil
}
