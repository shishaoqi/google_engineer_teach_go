package main

import (
	"fmt"
	)

type worker struct {
	in chan int
	done chan bool //通知完成
}

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d reveived %c\n", id, n)
		done <- true //通过 通信（channel） 来通知 -- 通知我接收完了
	}
}

func createWorker(chanId int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWorker(chanId, w.in, w.done)//// 此函数，我们可以 在此想做什么就做什么 (函数内加了 channel 来做阻塞控制)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'a' + i
		<-workers[i].done //在这收，收什么无所谓，反正我们接收进来后就意味着：doWorker操作完毕 （原理：doWorker中停止了block）
	}

	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'A' + i
		<-workers[i].done
	}
	//time.Sleep(time.Millisecond)//用 done chan 来同样实现相同效果
}

func main() {
	chanDemo()
}

