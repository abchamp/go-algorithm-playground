// happy number
package main

import "fmt"

// number repeat to 1 => happy

// example::
// input 19 => true
// 19 => 1 and 9 so 1^2 = 1, 9^9 = 81 to 1+81 => 82
// 82 => 8 and 2 so 8^2 = 64, 2^2 = 4 to 64+4 => 68
// 68 => 6 and 8 so 6^2 = 36, 8^2 = 64 to 36+64 => 100
// 1^2 => 2 , 0^2 => 0, 0^2 => 0, 1 is Happy

func isHappy(n int64) bool {
	var sum int64 = 0
	for n > 0 {
		var e int64 = n % 10 // 1234 => 4 (last)
		n = n / 10           // is breaking loop
		sum += e * e
		fmt.Println(sum, n)
	}

	// exist condition
	if sum == 1 {
		return true
	} else if (sum > 1) && (sum <= 4) {
		return false
	}
	//  recursive
	return isHappy(sum)
}

func main() {
	isHappy(19) // true
	isHappy(2)  // false
}
