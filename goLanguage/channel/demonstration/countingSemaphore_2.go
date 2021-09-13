package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int, seat Seat) {
	log.Print("customer#", c, " enters the bar")
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
	bar <- seat
}

// 计数信号量通常用于"强制"执行最大数量的并发请求.

// 通过发送获取所有权,通过接收释放 -- countingSemaphore_1.go 改版（有空位的情况下创建客户goroutine）
func main() {
	rand.Seed(time.Now().UnixNano())

	// the bar has 10 seats.
	bar24x7 := make(Bar, 10)
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block
		bar24x7 <- Seat(seatId)
	}
	/*for item := range bar24x7 {
		fmt.Println("bar24x7 channel: ", item)
	}*/
	fmt.Println(" ===== start server ......")

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// 创建出大量goroutine,消耗了资源
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)

		// need a seat to serve next customer
		seat := <- bar24x7 // 有空位的情况下创建客户goroutine
		go bar24x7.ServeCustomer(customerId, seat)

		log.Print("chan len: ", len(bar24x7))
	}

	for{time.Sleep(time.Second)} //避免程序退出
}
