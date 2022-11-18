package main

import "fmt"

func main() {
	x := 0

	// it references is known as closue

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

// use closure is by writing a function which return another function.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
