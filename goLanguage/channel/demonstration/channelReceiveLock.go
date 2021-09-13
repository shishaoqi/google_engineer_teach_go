package main

import "fmt"

// channel 用作互斥锁 -- 锁定接收示例

func main() {
	// The capacity must be one.
	mutex := make(chan struct{}, 1) // 注释掉 -- 会有奇效
	mutex <- struct{}{} // this line is needed ------ 为什么要有这行？为实现第一个接收有值（不锁定）

	counter := 0
	increase := func() {
		<-mutex // lock // 一起注释掉
		counter++
		mutex <- struct{}{} // unlock // 一起注释掉
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}

		done <- struct{}{}
	}

	d := make(chan struct{})
	go increase1000(d)
	go increase1000(d)
	go increase1000(d)
	go increase1000(d)
	<-d; <-d; <-d; <-d
	fmt.Println(counter)
}
