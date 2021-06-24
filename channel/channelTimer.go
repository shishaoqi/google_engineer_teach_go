package main

import (
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <- chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Println("Hi!")
	<- AfterDuration(time.Second)
	fmt.Println("Hello!")
	<- AfterDuration(time.Second)
	fmt.Println("Bye!")
}
// 事实上,After在功能time标准库提供相同的功能,具有更高效的实现.我们应该改用那个函数使代码看起来干净.
//
//请注意,<-time.After(aDuration)将使当前goroutine进入阻塞状态,但不会使time.Sleep(aDuration)函数调用.
//
//使用<-time.After(aDuration)通常用在下面将要引入的超时机制使用.
