package main

import "fmt"

func main() {
	bmis := []float64{1.2, 3.333, 4.5555}
	avgBmi := calculateAvg(bmis...)
	avgBmi = calculateAvgOnSlice(bmis)
	fmt.Println(avgBmi)
}

// 不定长输入参数
func calculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total
}

func calculateAvgOnSlice(bmis []float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total
}
