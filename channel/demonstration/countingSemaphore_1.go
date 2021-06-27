package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type Seat int // 坐位的编号
type SeatChan chan Seat // 这么去理解：酒巴有坐位，会轮流被客户占用。此坐位用 chan 来表示，类型取 Seat

type ServedCustomer struct  {
	n int
	mutex sync.Mutex
}

var scer ServedCustomer

func (sc SeatChan) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")
	seat := <- sc // need a seat to drink
	log.Print("++ customer#", c, " drinks at seat#", seat)
	serveTime := time.Second * time.Duration(2 + rand.Intn(30))
	time.Sleep(serveTime)
	log.Print("===== serviceTime：", serveTime)

	//log.Print("===== waiting customer number# ", c - ServedNum)
	sc <- seat
	scer.mutex.Lock()
	scer.n++
	scer.mutex.Unlock()
	log.Print("===== have serve number# ", scer.n)
}

// 计数信号量通常用于"强制"执行最大数量的并发请求.

// 通过发送获取所有权,通过接收释放
// 在上面的示例中,尽管在任何给定时间将有最多十个顾客在喝酒,但酒吧可能同时服务十个以上顾客.一些客户正在等待免费座位.
func main() {
	rand.Seed(time.Now().UnixNano())

	// the bar has 10 seats.
	bar24x7 := make(SeatChan, 10) // bar24x7 酒巴有10个坐位(即容量为 10)
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block
		bar24x7 <- Seat(seatId)
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// 创建出大量goroutine,消耗了资源
	for customerId := 0; ; customerId++ {
		log.Print("chan len: ", len(bar24x7))
		log.Print("waiting customer is ", customerId - scer.n - (10 - len(bar24x7)))
		go bar24x7.ServeCustomer(customerId)
		time.Sleep(time.Second)
	}

	for{time.Sleep(time.Second)} //避免程序退出
}
