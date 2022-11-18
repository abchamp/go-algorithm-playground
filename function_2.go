package main

import "fmt"

func main() {
	// fmt.Println(add(1, 2, 3))
	// var xs []int = []int{1, 2, 3}
	// fmt.Println(add(xs...))
	// ... mean {1, 2, 3} => (1, 2, 3)

	// Closure It is possible to create functions inside of function

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))
}

// Variadic Function

func add(args ...int) (n int) {
	// By using ... before the type name of the last parameter you can indicate
	// that it takes zero or more of those parameters
	// in this case we take zero or more ints we invoke the function
	// like any other other function except we can pass as many int as we want

	var total int = 0
	for _, v := range args {
		total += v
	}

	return total
}

// func Println(a ...interface{}) (n int, err error)
