package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	sub := Sub{0, 0}
	sub.executeFromScript("test-input.txt")

	if sub.HorizontalPosition != 15 {
		t.Errorf("Horizontal Position Received %v, Expected %v", sub.HorizontalPosition, 15)
	}

	if sub.Depth != 10 {
		t.Errorf("Depth Received %v, Expected %v", sub.Depth, 10)
	}
}

func TestPart2(t *testing.T) {

	sub := Sub2{0, 0, 0}
	sub.executeFromScript("test-input.txt")

	if sub.HorizontalPosition != 15 {
		t.Errorf("Horizontal Position Received %v, Expected %v", sub.HorizontalPosition, 15)
	}

	if sub.Depth != 60 {
		t.Errorf("Depth Received %v, Expected %v", sub.Depth, 60)
	}
}
