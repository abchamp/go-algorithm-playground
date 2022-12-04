package main

import (
	"fmt"
	"time"
)

func add(c chan int, a int, b int) {
	result := a + b
	c <- result
}

func multiply(c chan int) {
	result := <-c
	result *= 2

	fmt.Printf("Result is %d\n", result)
}

func main() {
	channel := make(chan int)
	a := 10
	b := 5

	// we want (a+b) * 2
	go add(channel, a, b)
	go multiply(channel)

	time.Sleep(1 * time.Second)
}
