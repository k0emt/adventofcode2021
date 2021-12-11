package main

import (
	"testing"
)

func TestFunction(t *testing.T) {

	mp := MetricParser{}

	loadData("test-input.dat", mp.addDataPoint, true)

	mp.counters()

	expected := 12
	if mp.state.counter != expected {
		t.Errorf("Expected to process %v lines, but got %v lines", expected, mp.state.counter)
	}

	expected = 5
	if mp.state.dataWidth != expected {
		t.Errorf("Expected data width %v, but got %v", expected, mp.state.dataWidth)
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
