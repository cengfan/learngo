package main

import (
	"fmt"
)

/**
创建通道并返回，同时开一个gorutine来异步接收数据
 */
func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(i, w.in, w.done)
	return w
}

type worker struct {
	in   chan int
	done chan bool
}

/**
先并发打印小写，然后并发打印大写
 */
func chainDemo() {

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for _, worker := range workers {
		<-worker.done
	}



	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done

	}

}

func doWork(i int, c chan int, d chan bool) {
	for n := range c {
		fmt.Printf("worker %d,received %c \n", i, n)

		d <- true

	}
}

func main() {
	chainDemo()

}
