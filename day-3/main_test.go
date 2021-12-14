package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	mp := MetricParser{}

	loadData("test-input.dat", mp.addDataPoint, true)

	mp.counters()

	expected := 12
	if mp.counter != expected {
		t.Errorf("Expected to process %v lines, but got %v lines", expected, mp.counter)
	}

	expected = 5
	if mp.dataWidth != expected {
		t.Errorf("Expected data width %v, but got %v", expected, mp.dataWidth)
	}

	mp.analyze()

	expected = 22
	if mp.Gamma != expected {
		t.Errorf("Gama Received %v, Expected %v", mp.Gamma, expected)
	}

	expected = 9
	if mp.Epsilon != expected {
		t.Errorf("Epsilon Received %v, Expected %v", mp.Epsilon, expected)
	}

	expected = 198
	if mp.Power != expected {
		t.Errorf("Power Received %v, Expected %v", mp.Power, expected)
	}
}

func TestPart2(t *testing.T) {

	mp := MetricParser{}

	loadData("test-input.dat", mp.addDataPoint, true)

	mp.counters()

	mp.calculateLifeSupport()

	expected := 23
	if mp.Oxygen != expected {
		t.Errorf("Oxygen Received %v, Expected %v", mp.Oxygen, expected)
	}

	expected = 10
	if mp.CO2 != expected {
		t.Errorf("CO2 Received %v, Expected %v", mp.CO2, expected)
	}

	expected = 230
	if mp.LifeSupport != expected {
		t.Errorf("Life Support Received %v, Expected %v", mp.LifeSupport, expected)
	}

}

func testNumbers() []string {
	numbers := []string{"00", "01", "11"}
	return numbers
}

func TestBitFrequencyPositionZero(t *testing.T) {

	zeroes, ones := determineBitFrequency(0, testNumbers())

	if zeroes != 2 {
		t.Errorf("expected 2 zeroes, got %v", zeroes)
	}

	if ones != 1 {
		t.Errorf("expected 1 one, got %v", ones)
	}
}

func TestBitFrequencyPositionOne(t *testing.T) {

	zeroes, ones := determineBitFrequency(1, testNumbers())

	if zeroes != 1 {
		t.Errorf("expected 1 zero, got %v", zeroes)
	}

	if ones != 2 {
		t.Errorf("expected 2 ones, got %v", ones)
	}
}

func TestBitCommon(t *testing.T) {

	leastCommon, mostCommon := determineCommon(0, testNumbers())

	if leastCommon != '1' {
		t.Errorf("least common should be 1, was: %v", leastCommon)
	}

	if mostCommon != '0' {
		t.Errorf("most common should be 0, was: %v", mostCommon)
	}
}

func TestBitCommonPos1(t *testing.T) {

	leastCommon, mostCommon := determineCommon(1, testNumbers())

	if leastCommon != '0' {
		t.Errorf("least common should be 0, was: %v", leastCommon)
	}

	if mostCommon != '1' {
		t.Errorf("most common should be 1, was: %v", mostCommon)
	}
}

func TestBitEqual(t *testing.T) {

	testNumbers := []string{"0", "1"}
	leastCommon, mostCommon := determineCommon(0, testNumbers)

	if leastCommon != '0' {
		t.Errorf("least common should be 0, was: %v", leastCommon)
	}

	if mostCommon != '1' {
		t.Errorf("most common should be 1, was: %v", mostCommon)
	}
}
