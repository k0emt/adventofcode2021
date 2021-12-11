package main

import (
	"fmt"
	"strconv"
)

type MetricParser struct {
	Gamma       int
	Epsilon     int
	Power       int
	state       internals
	zeroCounter [16]int
	oneCounter  [16]int
}

type internals struct {
	counter   int
	dataWidth int
}

func (mp *MetricParser) addDataPoint(dataPoint string, verbose bool) {
	if len(dataPoint) > mp.state.dataWidth {
		mp.state.dataWidth = len(dataPoint)
	}
	if verbose {
		fmt.Println(dataPoint)
	}
	mp.state.counter++
	for i := range dataPoint {
		if dataPoint[i] == '0' {
			mp.zeroCounter[i]++
		} else {
			mp.oneCounter[i]++
		}
	}
}

func (mp *MetricParser) analyze() {

	gamma := make([]byte, mp.state.dataWidth)
	epsilon := make([]byte, mp.state.dataWidth)

	for i := 0; i < mp.state.dataWidth; i++ {
		if mp.zeroCounter[i] > mp.oneCounter[i] {
			gamma[i] = '0'
			epsilon[i] = '1'
		} else {
			gamma[i] = '1'
			epsilon[i] = '0'
		}
	}

	tmp, _ := strconv.ParseInt(string(gamma), 2, 0)
	mp.Gamma = int(tmp)

	tmp, _ = strconv.ParseInt(string(epsilon), 2, 0)
	mp.Epsilon = int(tmp)

	mp.Power = mp.Gamma * mp.Epsilon
}

func (mp *MetricParser) counters() {
	fmt.Println("zeroes: ", mp.zeroCounter)
	fmt.Println("ones  : ", mp.oneCounter)
}
