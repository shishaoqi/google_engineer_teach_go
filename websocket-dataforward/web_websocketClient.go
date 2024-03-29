package main

import (
	"crypto/md5"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

const (
	TOKEN_SALT1 = "iCX0C0bj59xt"
	TOKEN_SALT2 = "g9C1Hr6O1WfR"
)

var host = "0.0.0.0"
var origin = "http://"+host+":9088/"

//var host = "test-cn-1-data.airdroid.com"
//var origin = "http://"+host+":9088/"
//var url = "ws://59.57.13.156:8088/phone?heartbeat=1&instance_id=1001&id=1001&token=3182d1050c65c7eeaa943c2ce893d843"

func main() {
	for i := 1; i <= 500; i++ {
		idStr := strconv.Itoa(i)
		tokenStr := GenerateToken(idStr)
		time.Sleep(10000 * time.Microsecond)
		go createClient(idStr, tokenStr)
	}

	time.Sleep(36000 * time.Second)
}

func createClient(id string, token string) {
	url := "ws://"+host+":8088/web?heartbeat=1&instance_id=%s&id=%s&token=%s"
	url = fmt.Sprintf(url, id, id, token)
	log.Println("===== request url =====:", id, url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal("connect fail: ", err)
	}
	message := []byte("client--("+id+") say hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Println("ws Write error: ", err)
		return
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	for {
		m, err := ws.Read(msg)
		if err != nil {
			log.Println("ws Read error: ", err)
			ws.Close()
			return
		}
		fmt.Printf("Receive: %s\n", msg[:m])
	}

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

