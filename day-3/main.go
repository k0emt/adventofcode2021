package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Advent of Code, Day 3")
	mp := MetricParser{0, 0, 0, make([]int, 5)}
	loadData("problem3.dat", mp.addDataPoint)
	fmt.Printf("Gamma: %v, Epsilon: %v, Power: %v\n", mp.Gamma, mp.Epsilon, mp.Power)
}

// execute from script
func loadData(dataFileName string, metricParserAddPoint func(string)) {
	file, err := os.Open(dataFileName)

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
		metricParserAddPoint(each_ln)
	}
}
