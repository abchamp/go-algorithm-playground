package main

import "fmt"

func main() {
	// var x [1]float64
	// fmt.Print(x[0])

	// 5 out of 10
	// x := make([]float64, 5, 10)

	// arr := [5]float64{1, 2, 3, 4, 5}
	// slices
	// x := arr[0:5]
	
	// append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 88, 99)
	fmt.Println(slice1, slice2)
	
	// copy
	slice1 = []int{1, 2, 3}
	slice2 = make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
