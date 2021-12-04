package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Advent of Code, Day 2")

	sub := Sub{0, 0}
	doScriptedNavigation("navigation-script.txt", sub.executeCommand)
	fmt.Printf("Depth: %d\nHorizontal Position: %d\nMultiplied: %d\n", sub.Depth, sub.HorizontalPosition, sub.Depth*sub.HorizontalPosition)

	sub2 := Sub2{0, 0, 0}
	doScriptedNavigation("navigation-script.txt", sub2.executeCommand)
	fmt.Printf("Depth: %d\nHorizontal Position: %d\nMultiplied: %d\n", sub2.Depth, sub2.HorizontalPosition, sub2.Depth*sub.HorizontalPosition)

}

// execute from script
func doScriptedNavigation(navigationScriptFileName string, subCommandRunner func(string)) {
	file, err := os.Open(navigationScriptFileName)

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
		subCommandRunner(each_ln)
	}
}
