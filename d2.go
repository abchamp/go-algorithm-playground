// Kadane' Algorithms & The Maximum Subarray Problem
// given an array of integers, return the maximum sum of a
// subarray of the array.
// [2, 1, -2, 3, 2]
//2 : [2], [2,1], [2,1,-2], [2, 1, -2, 3], [2, 1, -2, 3, 2]
//1 : [1]. [1. -2]. [1. -2, 3]..

// requirement
// continuous array => sum sub continuous array
package main

import "fmt"

func maxSubArray(inputN []int64) int64 {
	var maxCurrent int64 = inputN[0]
	var maxGlobal int64 = inputN[0]

	for i := 1; i < len(inputN); i++ {
		// compare between now value and maxCurrent + (now value)
		if inputN[i] > (maxCurrent + inputN[i]) {
			maxCurrent = inputN[i]
		} else {
			maxCurrent = maxCurrent + inputN[i]
		}

		// maxCurrent > max Global
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}

	return maxGlobal
}

func main() {
	var inputA []int64 = []int64{2, 1, -2, 3, 2}
	fmt.Println(maxSubArray(inputA))
}
