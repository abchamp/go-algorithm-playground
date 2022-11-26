package main

import (
	"fmt"
	"time"
)

func goroutine1(c chan bool) {
	fmt.Println("Goroutines #1 has started, waiting for Goroutines #2 to start")
	<-c // thread blocking read from channel and wait
	fmt.Println("Goroutines #1 received a notification from Goroutines #2")
}

func goroutine2(c chan bool) {
	fmt.Println("Goroutines #2 has started, do some work and notify Goroutines #1")
	time.Sleep(2 * time.Second) // simulate working
	c <- true                   // set
	fmt.Println("Goroutines #2 has finished")
}

func main() {
	c := make(chan bool)

	go goroutine1(c)
	go goroutine2(c)

	time.Sleep(3 * time.Second)
}

/**
func main() {
	c := make(chan bool)
	<- c

	fmt.Println("Done") // never reach this point
}
**/

/**
wait
Goroutines #1 has started, waiting for Goroutines #2 to start
Goroutines #2 has started, do some work and notify Goroutines #1/
..
Goroutines #2 has finished
Goroutines #1 received a notification from Goroutines #2

**/
