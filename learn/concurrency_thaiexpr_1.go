package main

import (
	"fmt"
	"strings"
	"time"
)

func processText(text string) {
	parts := strings.Split(text, "|")
	name := parts[0]
	age := parts[1]
	gender := parts[2]

	time.Sleep(3 * time.Second)
	fmt.Printf("Send email >> Name=%-5s gender=%-6s age=%2s\n", name, gender, age)
}

func main() {

	var allText []string = []string{
		"Perth|29|male",
		"Noom|25|male",
		"Onny|25|female",
	}

	for _, text := range allText {
		go processText(text)
		fmt.Println("Your request has been added to queue.")
	}

	time.Sleep(4 * time.Second)
	// Highlight: go func() {...}()
	// go func() {
	// 	for i := 0; i < 20; i++ {
	// 		time.Sleep(500 * time.Millisecond)
	// 		fmt.Println("Task 2")
	// 	}
	// }()

	// for i := 0; i < 20; i++ {
	// 	time.Sleep(500 * time.Millisecond)
	// 	fmt.Println("Task 1")
	// }
}
