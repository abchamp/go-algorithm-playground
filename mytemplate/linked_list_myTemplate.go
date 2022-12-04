package mytemplate

import "fmt"

type List struct {
	Head *MyNode
}

type MyNode struct {
	Value interface{}
	Next  *MyNode
}

func (l *List) Len() int {
	p := l.Head
	var i int = 0
	// add => add => nil
	for p.Next != nil {
		p = p.Next
		i++
		if p.Next == nil {
			i++
		}
	}
	return i
}

func (l *List) Push(value interface{}) {
	n := &MyNode{value, nil}
	if l.Head == nil {
		l.Head = n
	} else {
		w := l.Head
		for w.Next != nil {
			w = w.Next
		}
		w.Next = n
	}
}

func (l *List) Print() {
	p := l.Head
	for p != nil {
		fmt.Println(p.Value)
		p = p.Next
	}
}

// func main() {
// 	l := List{nil}
// 	l.Push(0)
// 	l.Push(9)
// 	l.Push(20)
// 	l.Print()
// }
