package main

import (
	"testing"
)

// QUESTION: is there a way to define TEST_NUMBERS as a reusable constant?
// currently Go only supports basic data types for constants

func testNumbers() []int {
	numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	return numbers
}

func TestLargerNumbersExample(t *testing.T) {

	actual := LargerNumbers(testNumbers())
	expected := 7

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSlidingWindowLargerNumbersExample(t *testing.T) {

	actual := SlidingWindowLargerNumbers(testNumbers())
	expected := 5

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
