// Finding the Only Single Number in an Array
// [2, 1, 4, 4, 2] => output: 1, 2, and 4
package main

import (
	"fmt"
	"sort"
)

func main() {
	var x []int = []int{2, 1, 1, 4, 8, 4, 2}
	var y int = solutionOne(x)
	fmt.Println(y)
}

func solutionOne(inp []int) int {
	sort.Ints(inp)
	fmt.Println("Inp:", inp)
	for i := 0; i < len(inp); i++ {
		if i == 0 {
			if inp[i] != inp[i+1] {
				return inp[i]
			}
		} else if i == len(inp)-1 {
			if inp[i] != inp[i-1] {
				return inp[i]
			}
		} else {
			if inp[i] != inp[i-1] && inp[i] != inp[i+1] {
				return inp[i]
			}
		}
	}
	return -1
	// var strs []string = []string{"c", "a", "b"}
	// sort.Strings(strs)
	// fmt.Println("Strings", strs)

	// var ints []int = []int{7, 2, 4}
	// sort.Ints(ints)
	// fmt.Println("Ints: ", ints)

	// var s = sort.IntsAreSorted(ints)
	// fmt.Println("Sorted: ", s)
}
