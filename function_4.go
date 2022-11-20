package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC 555555")

}
