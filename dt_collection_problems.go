package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var min_x int = x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < min_x {
			min_x = x[i]
		}
	}
	fmt.Println("smallest number is ", min_x)
}
