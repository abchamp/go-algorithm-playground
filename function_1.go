package main

import "fmt"

func main() {
	// var xs []float64 = []float64{98, 93, 77, 82, 83}

	// var total float64 = 0.0

	// for _, v := range xs {
	// 	total += v
	// }

	// fmt.Println(total / float64(len(xs)))

	var xs []float64 = []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	// the names of the parameter don't have to match in ther calling function
}

func returnMultipleValue() (int, int) {
	return 5, 6
}

func average(xs []float64) float64 {
	// panic("Not Implemented")
	var total float64 = 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}
