package main

import "fmt"

type MetricParser struct {
	Power          int
	Gamma          int
	Epsilon        int
	dataPointCount []int
}

func (mp *MetricParser) addDataPoint(dataPoint string) {
	fmt.Println(dataPoint)
}
