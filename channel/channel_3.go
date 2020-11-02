package main

import (
	"fmt"
	"time"
)

func createWorker(chanId int) chan<- int { // 接收到此chan的人，只能向这chan里发数据
	c := make(chan int)

	// 此处，我们可以把此匿名函数 从 函数参数里传进来，从而实现想做什么就让客户传什么进来（）
	go func() {
		for {
			fmt.Printf("Worker'id %d received %c\n", chanId, <-c) // 此处从 chan 收数据
		}
	}()

	return c
}

func chanDemo3(stop chan bool) {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i) // createWorker
	}

	/*go func() { //测试这边能放吗
		for i := 0; i < 9; i++{
			fmt.Printf("Worker received %d\n", <-channels[i])
		}
	}()*/

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	go func() {
		fmt.Println("aaaaaaaaa")
		for {
			if <-stop {
				break
			}
			fmt.Println("111111111")
		}
		fmt.Println("bbbbbbbbb")
	}()
}

func main() {
	stop := make(chan bool)
	chanDemo3(stop)
	stop <- true
	time.Sleep(time.Millisecond)
}
