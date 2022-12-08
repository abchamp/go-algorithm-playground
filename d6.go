// Solve Stock Span Problem
package main

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		lastIndex := len(*s) - 1
		element := (*s)[lastIndex]
		*s = (*s)[:lastIndex]
		return element, true
	}
}

func (s *Stack) Peak() int {
	if s.IsEmpty() {
		return -1
	} else {
		element := (*s)[len(*s)-1]
		return element
	}
}

func solveProblem(stack Stack) {
	fmt.Println(stack.Peak())

}

func main() {
	var stack Stack
	// solveProblem(&stack)
}
