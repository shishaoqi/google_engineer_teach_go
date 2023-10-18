package main

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"net"
)

type Ws struct {
	Id   int
	Conn net.Conn
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	ws := new(Ws)
	ws.Id = 100999
	ws.Conn = CreateClient("localhost:9100")

	ws.Reveive()
	ws.Register()
	ws.Heart()

	wg.Wait()
}

func CreateClient(url string) net.Conn {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal("connect fail: ", err)
	}

	return conn
}

// 注册
func (ws Ws) Register() {
	ws.Conn.Write([]byte(`#*#*head*#*#{"register":"` + "1000" + `","secretKey":"aaa"}#*#*end*#*#`))
}

// 监听返回消息
func (ws Ws) Reveive() {
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := ws.Conn.Read(buf)
			if err == io.EOF {
				log.Println("connect err:", err)
				return
			}
			log.Println("收到的消息: ", string(buf[0:n]))
		}
	}()
}

// 心跳
func (ws Ws) Heart() {
	go func() {
		msg := fmt.Sprintf("{\"heart\":%v,\"secretKey\":\"workMan_client_id_%v\"}", ws.Id, ws.Id)

		for true {
			ws.Conn.Write([]byte(msg))
			time.Sleep(time.Second * 30)
		}
	}()
}
