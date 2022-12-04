package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool, 3)

	for true {
		go func() {
			fmt.Println(time.Now().Second())
			time.Sleep(10 * time.Second)
			<-c
		}()
		fmt.Print("Hello \n")
		c <- true
	}
}

/**
call
Hello
Hello
Hello
Hello
from thred: 29 <= 10
from thred: 29 <= 10
from thred: 29 <= 10
2from thred: 9 <= 10
**/
