package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator1() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker1(id int, c chan int) {
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

func createWorker1(chanId int) chan<- int {
	c := make(chan int)
	go worker1(chanId, c)
	return c
}

func main() {
	var c1, c2 = generator1(), generator1() // c1 and c2 = nil
	w := createWorker1(0)

	for {
		//有缺点
		select {
		case n := <-c1: //收到一个数，（后面做的事）又会阻塞住
			w <- n //后面做的事
		case n := <-c2:
			w <- n
		}
	}
}
