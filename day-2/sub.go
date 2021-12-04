package main

import (
	"bufio"
	"log"
	"os"
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

// execute from script
func (s *Sub) executeFromScript(fileName string) {
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
