package main

import (
	"fmt"
	"sync"
)

/**
创建通道并返回，同时开一个gorutine来异步接收数据
 */
func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(i, w)
	return w
}

type worker struct {
	in   chan int
	done func()
}

/**
并发打印大小写,使用sync.WaitGroup来控制
 */
func chainDemo() {

	var workers [10]worker

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}



}

func doWork(i int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d,received %c \n", i, n)

		w.done()

	}
}

func main() {
	chainDemo()

}
