package main

import (
	"strconv"
	"strings"
)

type Sub struct {
	Depth              int
	HorizontalPosition int
}

func (s *Sub) goForward(increment int) {
	s.HorizontalPosition += increment
}

func (s *Sub) goDown(increment int) {
	s.Depth += increment
}

func (s *Sub) goUp(increment int) {
	s.Depth -= increment
}

func (s *Sub) executeCommand(command string) {

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
