package main

import (
	"fmt"
)

func main() {

	fmt.Println("Advent of Code, Day 2")

	sub := Sub{0, 0}
	sub.executeFromScript("navigation-script.txt")
	fmt.Printf("Depth: %d\nHorizontal Position: %d\nMultiplied: %d\n", sub.Depth, sub.HorizontalPosition, sub.Depth*sub.HorizontalPosition)

	sub2 := Sub2{0, 0, 0}
	sub2.executeFromScript("navigation-script.txt")
	fmt.Printf("Depth: %d\nHorizontal Position: %d\nMultiplied: %d\n", sub2.Depth, sub2.HorizontalPosition, sub2.Depth*sub.HorizontalPosition)

}
