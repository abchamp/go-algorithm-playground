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

func solveProblem(stack Stack, prices []int) {
	// fmt.Println(stack.Peak())
	//
	// var currentValue int = prices[0]
	// var span []int // slice declaration; no memory allocation
	var span = make([]int, 0)
	//slice initialization with length (0) and capacity (10);
	//memory allocated for 10 ints
	stack.Push(0)

	for i := 0; i < len(prices); i++ {
		span = append(span, 0)
	}

	span[0] = 1

	for i := 0; i < len(prices); i++ {
		for {
			if !stack.IsEmpty() {
				if prices[stack.Peak()] <= prices[i] {
					stack.Pop()
				} else {
					break
				}
			} else {
				break
			}
		}
		//

		if stack.IsEmpty() {
			// Case: 1
			span[i] = i + 1
		} else {
			// Case: 2
			span[i] = i - stack.Peak()
		}

		stack.Push(i)
	}

	fmt.Println(span)
}

func main() {
	var stack Stack
	var prices []int = []int{90, 40, 20, 30, 80, 60, 100}
	// &stack
	solveProblem(stack, prices)
}
