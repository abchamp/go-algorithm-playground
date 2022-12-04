// find the middle of link list
package main

import (
	"abchamp/myalgor/mytemplate"
	"fmt"
)

type myList struct {
	head *mytemplate.MyNode
}

func (l *myList) Push(value interface{}) {
	n := &mytemplate.MyNode{value, nil}
	if l.head == nil {
		l.head = n
	} else {
		w := l.head
		for w.Next != nil {
			w = w.Next
		}
		w.Next = n
	}
}

func (l *myList) Print() {
	p := l.head
	for p != nil {
		fmt.Println(p.Value)
		p = p.Next
	}
}

func (l *myList) FindMiddle() {
	fast := l.head
	slow := l.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Print(slow.Value)
}

func main() {
	l := myList{nil}
	for i := 1; i <= 9; i++ {
		l.Push(i)
	}

	// l.Push(2)
	// l.Push(3)
	// l.Push(4)
	// l.Push(5)
	l.FindMiddle()
}
