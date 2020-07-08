package main

import (
	"fmt"
	"github.com/shishao/hello/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T ---- %v\n", r, r)

	//fmt.Println(download(r))
}