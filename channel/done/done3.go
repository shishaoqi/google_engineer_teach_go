//大小写分开打
package main

import (
	"fmt"
	)

func doWorker3(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d reveived %c\n", id, n)

		//版本2的第一版本；
		done <- true //（所有 channel的发都是阻塞式的，给 done 发一个，另一端必须要有一接收方）
		//worker.in <- 'a' + i 接收完后，在等待另端接收过程时，worker.in <- 'A' + i 又发送来，于是 fatal error: all goroutines are asleep - deadlock!

		//版本2的第二版本；-- 大小写分开打，版本1就又可以了
		//go func(){ done <- true }()
	}
}

type worker3 struct {
	in chan int
	done chan bool
}

func createWorker3(chanId int) worker3 {
	w := worker3{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWorker3(chanId, w.in, w.done)
	return w
}

func chanDemo3() {
	var workers [10]worker3
	for i := 0; i < 10; i++ {
		workers[i] = createWorker3(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done //在这收，收什么无所谓，反正我们收进来后就意味着：doWorker操作完毕
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done //在这收，收什么无所谓，反正我们收进来后就意味着：doWorker操作完毕
	}
}

func main() {
	chanDemo3()
}

