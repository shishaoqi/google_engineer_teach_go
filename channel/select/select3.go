package main

import (
	"fmt"
	"time"
	"math/rand"
)

// 数据源的生成
func generator3() chan int{
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

func worker3(id int, c chan int) {
	for n := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d reveived %d\n", id, n)
	}

	/*for n, ok := <-c {
		if !ok {
			break
		}
		fmt.Printf("Worker %d reveived %d\n", id, n)
	}*/
}

func createWorker3(chanId int) chan<- int {
	c := make(chan int)
	go worker3(chanId, c)
	return c
}

func main() {
	var c1, c2 = generator3(), generator3() // c1 and c2 = nil
	var worker  = createWorker3(0)

	var values []int
	tm := time.After(10 * time.Second) //10秒钟后，结束for循环
	tick := time.Tick(2 * time.Second) //Tick: 每隔这段时间返回一个channel

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		// 定时器的使用
		case <-time.After(800 * time.Millisecond): //每两个select的case之间的时间间隔
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <- tm: //(整个for循环执行的时间)10秒钟后，结束for循环
			fmt.Println("bye")
			return
		}
	}
}
