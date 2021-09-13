package main

import (
	"fmt"
	"time"
)

func worker2(chanId int, c chan int) {
	for {
		n := <- c
		fmt.Printf("Worker'id %d received %c\n", chanId, n)
	}
}

func chanDemo2() {
	var channels [10]chan int //
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) // make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。
		go worker2(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo2()
}
