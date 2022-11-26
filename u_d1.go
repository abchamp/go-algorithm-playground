// reverse an integer
package main

import (
	"fmt"
	"math"
)

func reverseNum(num int64) int64 {
	// check postive number
	var pos bool = true
	if num < 0 {
		pos = false
		num = int64(math.Abs(float64(num)))
	}
	//
	var result int64 = 0
	var mod int64 = 0

	for num > 0 {
		mod = num % 10
		result = (result * 10) + mod
		num = num / 10
	}

	if !pos {
		result = -1 * result
	}

	return result
}

func main() {
	result := reverseNum(-12345)
	fmt.Println(result)
	// var x int = 10
	// y := float64(x)
	// fmt.Println(math.Floor(y))
}
