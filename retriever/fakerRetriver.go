package main

import (
	"fmt"
	"github.com/shishao/hello/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值
	fmt.Println(download(r))
}