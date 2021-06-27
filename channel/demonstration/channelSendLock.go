package main

import "fmt"

// channel 用作互斥锁 -- 锁定发送示例

func main() {
	// The capacity must be one.
	mutex := make(chan struct{}, 1) // 注释掉 -- 会有奇效

	counter := 0
	increase := func() {
		mutex <- struct{}{} // lock // 注释掉
		counter++
		<-mutex // unlock // 注释掉
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
