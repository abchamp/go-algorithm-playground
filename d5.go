// Backspace String Comparisons
package main

import (
	"fmt"
	"strings"
)

// input ab#c and ad#c
// # is a backspace character, which means it deletes the previous character in the string
// return whether or not the two string are equal
// queue and indexer
func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func popqueue(queue []string) []string {
	return queue[0 : len(queue)-1]
}

func main() {
	var queue []string
	var input string = "ab#c"
	for i := 0; i < len(input); i++ {
		// fmt.Printf("%c\n", input[i])
		var inputAlphabet string = string(input[i])
		if strings.Compare(inputAlphabet, "#") == 0 {
			queue = popqueue(queue)
		} else {
			queue = enqueue(queue, inputAlphabet)
		}
	}

	fmt.Println(queue)
}
