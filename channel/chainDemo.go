package main

import (
	"fmt"
	"time"
)

func chanDemo() {

	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2

	time.Sleep(time.Microsecond)
}
func main() {
	chanDemo()
}
