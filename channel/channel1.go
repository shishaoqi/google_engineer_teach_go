package main

import (
	"fmt"
	"time"
)

func worker( c chan int) {
	for {
		n := <- c
		fmt.Println(n)
	}
}

func chanDemo() {
	//var c chan int // 这样初始化的c为： c == nil
	c := make(chan int)
	go worker(c) // goroutine
	c <- 1
	c <- 2

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
