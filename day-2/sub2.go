package main

import (
	"bufio"
	"log"
	"os"
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

// NOTE: seems like there should be an elegant way to remove this duplication between subs
func (s *Sub2) executeFromScript(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, each_ln := range text {
		// fmt.Println(each_ln)
		s.executeCommand(each_ln)
	}
}
