package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readFromChannel(channel chan int) {
	for number := range channel {
		fmt.Printf("Got '%d' from channel\n", number)
	}
}

func main() {

	// Note:: Read from Loop
	// channel := make(chan int)
	// go readFromChannel(channel)

	// for true {
	// 	time.Sleep(1 * time.Second)
	// 	channel <- rand.Intn(100)
	// }

	channel := make(chan bool)
	rand.Seed(time.Now().Unix()) // set up seed for random

	go func() {
		// ตัวนี้อาจจะออกมาก่อน
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // simulate working
		channel <- true
	}()
	// อาจจะเป็นตัวนี้ก็ได้
	timeout := time.Tick(3 * time.Second) // timeout = 3 seconds

	select {
	case <-timeout:
		fmt.Println("Timeout!")
	case <-channel:
		fmt.Println("Work done!")
	}

}
