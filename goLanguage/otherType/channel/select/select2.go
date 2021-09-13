package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator2() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			//time.Sleep(time.Duration(5 * time.Millisecond)//睡眠时间设置长点，让channel中的数据被冲刷掉
			out <- i
			i++
		}
	}()
	return out
}

func worker2(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d reveived %d\n", id, n)
	}
	/*for n, ok := <-c {
		if !ok {
			break
		}
		fmt.Printf("Worker %d reveived %d\n", id, n)
	}*/
}

func createWorker2(chanId int) chan<- int {
	c := make(chan int)
	go worker2(chanId, c)
	return c
}

func main() {
	var c1, c2 = generator2(), generator2() // c1 and c2 = nil
	var worker  = createWorker2(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int //利用 刚创建chan 为 nil, 就不会被select到
		if hasValue { //一轮循环中，判断接收方有值，那么worker赋值给activeWorker,从而能在select中能被执行（即执行worker2()）
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}
