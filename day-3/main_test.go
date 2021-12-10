package main

import (
	"testing"
)

func TestFunction(t *testing.T) {

	mp := MetricParser{0, 0, 0, make([]int, 5)}

	loadData("test-input.dat", mp.addDataPoint)

	if mp.Gamma != 22 {
		t.Errorf("Gama Received %v, Expected %v", mp.Gamma, 22)
	}

	if mp.Epsilon != 9 {
		t.Errorf("Epsilon Received %v, Expected %v", mp.Epsilon, 9)
	}

	if mp.Power != 198 {
		t.Errorf("Power Received %v, Expected %v", mp.Power, 190)
	}
}
