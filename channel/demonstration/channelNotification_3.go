package main

import (
	"log"
	"time"
)

// 对 channelNotification_1 与 channelNotification_2 用例扩展，实现 N对1 和 1对N通知

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) { // 左接右发：<-chan 表示:从chan接收数据， chan<- 表示：向chan发数据
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// Simulate an initailization phase.
	time.Sleep(time.Second * 3 / 2)
	// 1 to N notification.
	ready <- T{}; ready <- T{}; ready <- T{}
	// Being N-to-1 notified
	<-done; <-done; <-done

	//实际上,在本小节中介绍的进行1-to-N和N-to-1通知的方法在实践中并不常用.
	//在实践中,我们经常使用 sync.WaitGroup N对1的通知,而我们通过封闭渠道进行1对N的通知.请阅读下一部分以了解详细信息.
}
