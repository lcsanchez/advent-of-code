package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Submarine struct {
	Aim      int
	Position *Point
}

type Command func(*Submarine)

func adjustAim(i int) Command {
	return func(s *Submarine) {
		s.Aim += i
	}
}

func move(i int) Command {
	return func(s *Submarine) {
		s.Position.X += i
		s.Position.Y += i * s.Aim
	}
}

func readCommands(r io.Reader) ([]Command, error) {
	commands := make([]Command, 0)
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
			commands = append(commands, move(magnitude))
		case "down":
			commands = append(commands, adjustAim(magnitude))
		case "up":
			commands = append(commands, adjustAim(-magnitude))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return commands, nil
}

func navigateSubmarine(s *Submarine, commands []Command) {
	for _, command := range commands {
		command(s)
	}
}
