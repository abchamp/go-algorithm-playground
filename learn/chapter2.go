package main

import "fmt"

// this is a comment

func main() {
	/**
	fmt.Println("Hello World")
	fmt.Println("1 + 1 =", 1 + 1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(32132 * 42452)
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	**/
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
