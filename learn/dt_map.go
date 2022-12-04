package main

import "fmt"

func main() {
	// x  := make(map[string]int)
	// x["key"] = 10
	// fmt.Println(x["key"])

	x := make(map[int]int)
	fmt.Println(len(x))
	x[1] = 10
	fmt.Println(x[1])
	fmt.Println(len(x))

}
