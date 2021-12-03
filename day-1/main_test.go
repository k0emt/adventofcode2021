package main

import (
	"testing"
)

// QUESTION: is there a way to define TEST_NUMBERS as a reusable constant?

func TestLargerNumbersExample(t *testing.T) {

	actual := LargerNumbers([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	expected := 7

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSlidingWindowLargerNumbersExample(t *testing.T) {

	actual := SlidingWindowLargerNumbers([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	expected := 5

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
