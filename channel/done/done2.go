//大小写一起(同时)打
package main

import (
	"fmt"
	)

func doWorker2(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d reveived %c\n", id, n)

		// 版本2的第一版本；
		// done <- true //（所有 channel的发都是阻塞式的，给 done 发一个，另一端必须要有一接收方）
		// worker.in <- 'a' + i 接收完后，在等待另端接收过程时，worker.in <- 'A' + i 又发送来，于是 fatal error: all goroutines are asleep - deadlock!
		// 第二版本 通过 再创建个 goroutine 来解决 deadlock

		//版本2的第二版本；
		go func(){ done <- true }()
	}
}

type worker2 struct {
	in chan int
	done chan bool
}

func createWorker2(chanId int) worker2 {
	w := worker2{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWorker2(chanId, w.in, w.done)
	return w
}

func chanDemo2() {
	var workers [10]worker2
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i)
	}

	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done //在这收，收什么无所谓，反正我们收进来后就意味着：doWorker操作完毕
		<-worker.done //在这收，收什么无所谓，反正我们收进来后就意味着：doWorker操作完毕
	}
}

func main() {
	chanDemo2()
}

