package main

import (
	"crypto/md5"
	"fmt"
	"code.google.com/p/go.net/websocket"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	TOKEN_SALT1 = "iCX0C0bj59xt"
	TOKEN_SALT2 = "g9C1Hr6O1WfR"
)

var host = "test-cn-1-data.airdroid.com"
var origin = "http://"+host+":9088/"
//var url = "ws://59.57.13.156:8088/phone?heartbeat=1&instance_id=1001&id=1001&token=3182d1050c65c7eeaa943c2ce893d843"

func main() {
	for i := 1; i <= 3; i++ {
		time.Sleep(300 * time.Millisecond)
		idStr := strconv.Itoa(i)
		tokenStr := GenerateToken(idStr)
		go createClient(idStr, tokenStr)
	}

	time.Sleep(36000 * time.Second)
}

func createClient(id string, token string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("==== recover ====", r)
		}
	}()

	n := rand.Intn(50)
	time.Sleep(time.Duration(n) * 1000 * time.Microsecond)

	url := "ws://"+host+":8088/phone?heartbeat=1&instance_id=%s&id=%s&token=%s&device_type=21"
	url = fmt.Sprintf(url, id, id, token)
	log.Println("===== request url =====:", id, url)

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Println("connect fail: ", err)
		return
	}
	message := []byte("hello, world!你好")
	_, err = ws.Write(message)
	if err != nil {
		log.Println("ws write error: ", err)
		return
	}
	fmt.Printf("Success! Send: %s\n", message)

	go func(){
		var msg = make([]byte, 512)
		for {
			m, err := ws.Read(msg)
			if err != nil {
				log.Println("ws read error: ", err)
				ws.Close()
				return
			}
			fmt.Printf("Receive: %s\n", msg[:m])
		}
	}()

	t1 := time.NewTimer(40 * time.Second)
	//defer t1.Stop()
	deviceInfoTimer := time.NewTimer(90 * time.Second)
	//defer deviceInfoTimer.Stop()
	go func(t *time.Timer, id string, ws *websocket.Conn){

		for {
			select {
				case <- t.C:
					fmt.Print("h heartbeat -- " + time.Now().Format("2006-01-02 15:04:05") + " -- ")
					message := []byte("h")
					_, err = ws.Write(message)
					if err != nil {
						log.Println("ws write error: ", err)
						return
					}
					t.Reset(40 * time.Second)

				case <-deviceInfoTimer.C:
					fmt.Print("h+ heartbeatDeviceInfo -- " + time.Now().Format("2006-01-02 15:04:05") + " -- ")
					msg := []byte("h;;;1;;;1;;;{\"account_id\": 20132, \"device_id\": \"xx\", \"wifi\": 1 ,\"gsm\":2 ,\"battery\":95 ,\"port\":5457,\"ip\":\"127.0.0.1\"")
					_, err = ws.Write(msg)
					if err != nil {
						log.Println("ws write error: ", err)
						return
					}
					deviceInfoTimer.Reset(90 * time.Second)
			}
		}
	}(t1, id, ws)

	//ws.Close()//关闭连接
}

func GenerateToken(device_id string) string {
	return getMD5(getMD5(device_id+TOKEN_SALT1) + TOKEN_SALT2)
}

func getMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

