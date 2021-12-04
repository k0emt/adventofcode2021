package main

import (
	"strconv"
	"strings"
)

type Sub2 struct {
	Depth              int
	HorizontalPosition int
	Aim                int
}

func (s *Sub2) goForward(increment int) {
	s.HorizontalPosition += increment
	s.Depth += (s.Aim * increment)
}

func (s *Sub2) goDown(increment int) {
	s.Aim += increment
}

func (s *Sub2) goUp(increment int) {
	s.Aim -= increment
}

func (s *Sub2) executeCommand(command string) {
	cmd := strings.Split(command, " ")
	directive := cmd[0]
	amount, _ := strconv.Atoi(cmd[1])

	switch directive {
	case "forward":
		s.goForward(amount)
	case "down":
		s.goDown(amount)
	case "up":
		s.goUp(amount)
	}
}
