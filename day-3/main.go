package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Advent of Code, Day 3")
	mp := MetricParser{}
	loadData("problem3.dat", mp.addDataPoint, false)
	mp.analyze()
	fmt.Printf("Gamma: %v, Epsilon: %v, Power: %v\n", mp.Gamma, mp.Epsilon, mp.Power)

	mp.calculateLifeSupport()
	fmt.Printf("O2: %v, CO2: %v, LS: %v\n", mp.Oxygen, mp.CO2, mp.LifeSupport)
}

// execute from script
func loadData(dataFileName string, metricParserAddPoint func(string, bool), verbose bool) {
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
		metricParserAddPoint(each_ln, verbose)
	}
}
