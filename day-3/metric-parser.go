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

	LifeSupport int
	Oxygen      int
	CO2         int
}

type internals struct {
	counter    int
	dataWidth  int
	dataPoints []string
}

func (mp *MetricParser) addDataPoint(dataPoint string, verbose bool) {

	if verbose {
		fmt.Println(dataPoint)
	}

	if len(dataPoint) > mp.state.dataWidth {
		mp.state.dataWidth = len(dataPoint)
	}

	mp.state.dataPoints = append(mp.state.dataPoints, dataPoint)

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

func (mp *MetricParser) calculateLifeSupport() {

	mp.Oxygen = calculateOxygen(0, mp.state.dataPoints)

	mp.CO2 = calculateCO2(0, mp.state.dataPoints)

	mp.LifeSupport = mp.Oxygen * mp.CO2
}

func calculateOxygen(bitPosition int, dataRows []string) int {
	retVal := 0

	if len(dataRows) > 1 {
		_, most := determineCommon(bitPosition, dataRows)
		filteredDataRows := []string{}

		for i := range dataRows {
			if string(dataRows[i][bitPosition]) == string(most) {
				filteredDataRows = append(filteredDataRows, dataRows[i])
			}
		}

		return calculateOxygen(bitPosition+1, filteredDataRows)

	} else {
		tmp, _ := strconv.ParseInt(string(dataRows[0]), 2, 0)
		retVal = int(tmp)
	}

	return retVal
}

func calculateCO2(bitPosition int, dataRows []string) int {
	retVal := 0

	if len(dataRows) > 1 {
		least, _ := determineCommon(bitPosition, dataRows)
		filteredDataRows := []string{}

		for i := range dataRows {
			if string(dataRows[i][bitPosition]) == string(least) {
				filteredDataRows = append(filteredDataRows, dataRows[i])
			}
		}

		return calculateCO2(bitPosition+1, filteredDataRows)

	} else {
		tmp, _ := strconv.ParseInt(string(dataRows[0]), 2, 0)
		retVal = int(tmp)
	}

	return retVal
}

func determineBitFrequency(position int, readings []string) (int, int) {
	zeroes, ones := 0, 0

	for i := range readings {
		if readings[i][position] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	return zeroes, ones
}

func determineCommon(position int, readings []string) (leastCommon rune, mostCommon rune) {
	zeroes, ones := determineBitFrequency(position, readings)

	if zeroes > ones {
		leastCommon = '1'
		mostCommon = '0'
	} else {
		leastCommon = '0'
		mostCommon = '1'
	}

	return leastCommon, mostCommon
}

func (mp *MetricParser) counters() {
	fmt.Println("zeroes: ", mp.zeroCounter)
	fmt.Println("ones  : ", mp.oneCounter)
}
