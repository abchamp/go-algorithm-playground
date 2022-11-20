package main

import "fmt"

// func zero(x int) {
// 	// argument is copied to the function
// 	x = 0
// }

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	// basic used
	// var x int = 5
	// zero(&x)
	// fmt.Println(x)

	// another way to get a pointer is to use the built-in new function
	// new takes a type as an argument, allocates enough memory to fit a value of that type
	// and returns a pointer to it.

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

}
