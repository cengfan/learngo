package main

import (
	"fmt"
	"time"
)

/**
创建通道并返回，同时开一个gorutine来异步接收数据
 */
func createWorker(i int) chan<- int {
	c := make(chan int)
	go work(i,c)
	return c
}

func chainDemo() {
	var channels [10] chan<- int
//并行10个
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	//等待打印
	time.Sleep(time.Millisecond)
}

func work(i int,c chan int)  {
	for n:=range c{

		fmt.Printf("worker %d,received %c \n", i, n)
	}
}

func bufferedChannel(){
	c :=make(chan int,3)
	go work(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}
func channelClose()  {
	c :=make(chan int)
	go work(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chainDemo()

	//bufferedChannel()
	//channelClose()
}
