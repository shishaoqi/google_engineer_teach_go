package main

import (
	"fmt"
	"time"
)

func createWorker4(chanId int) chan<- int {
	c := make(chan int)
	go worker4(chanId, c)
	return c
}

func worker4(id int, c chan int) {
	/*for {
		// 判断 channel 已经关闭
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker'id %d received %c\n", id, n)
	}*/
	for n := range c { // 判断 channel 已经关闭 -- (优化方式)
		fmt.Printf("Worker %d reveived %c\n", id, n)
	}
}

func chanDemo4() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker4(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //缓冲区大小为 3
	go worker4(0, c)

	c <- 'a'
	c <- 'b'
	//c <- 'c'
	//c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker4(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo4()

	fmt.Println("\nBuffered channel")
	bufferedChannel()

	fmt.Println("\nChannel close and range")
	channelClose()
}
