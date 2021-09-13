package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest2(r chan<- int32)  {
	// 模拟耗时任务
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares2(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest2(ra)
	go longTimeRequest2(rb)

	fmt.Println(sumSquares2(<-ra, <-rb))
}
