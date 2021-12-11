package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var inputBytes []byte

type Point struct {
	X int
	Y int
}

type Submarine struct {
	Aim      int
	Position *Point
}

type CommandBuilderFunc func(direction string, magnitude int) (Command, error)
type Command func(*Submarine)

func MoveVertical(i int) Command {
	return func(s *Submarine) {
		s.Position.Y += i
	}
}

func MoveHorizontal(i int) Command {
	return func(s *Submarine) {
		s.Position.X += i
	}
}

func AdjustAim(i int) Command {
	return func(s *Submarine) {
		s.Aim += i
	}
}

func MoveWithAim(i int) Command {
	return func(s *Submarine) {
		s.Position.X += i
		s.Position.Y += i * s.Aim
	}
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
	submarine := &Submarine{
		Position: &Point{},
	}

	err := navigateSubmarine(submarine, bytes.NewReader(inputBytes), SubmarineCommandBuilder)
	if err != nil {
		return 0, err
	}

	return submarine.Position.X * submarine.Position.Y, nil
}

func partTwo() (int, error) {
	submarine := &Submarine{
		Position: &Point{},
	}

	err := navigateSubmarine(submarine, bytes.NewReader(inputBytes), SubmarineWithAimCommandBuilder)
	if err != nil {
		return 0, err
	}

	return submarine.Position.X * submarine.Position.Y, nil
}

func SubmarineCommandBuilder(direction string, magnitude int) (Command, error) {
	switch direction {
	case "forward":
		return MoveHorizontal(magnitude), nil
	case "down":
		return MoveVertical(magnitude), nil
	case "up":
		return MoveVertical(-magnitude), nil
	default:
		return nil, errors.New("unrecognized command")
	}
}

func SubmarineWithAimCommandBuilder(direction string, magnitude int) (Command, error) {
	switch direction {
	case "forward":
		return MoveWithAim(magnitude), nil
	case "down":
		return AdjustAim(magnitude), nil
	case "up":
		return AdjustAim(-magnitude), nil
	default:
		return nil, errors.New("unrecognized command")
	}
}

func navigateSubmarine(sub *Submarine, r io.Reader, f CommandBuilderFunc) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()

		parts := strings.Split(s, " ")
		if len(parts) != 2 {
			return fmt.Errorf("invalid command '%s'", s)
		}

		direction := parts[0]
		magnitude, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid magnitude '%s'", parts[1])
		}

		cmd, err := f(direction, magnitude)
		if err != nil {
			return err
		}
		cmd(sub)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
