package main

import (
	"time"
	"fmt"
)

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)
	c <- 1
	c <- 2
	n := <-c
	fmt.Println(n)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
