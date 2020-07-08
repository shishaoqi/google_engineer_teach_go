package main

import (
	"fmt"
	"github.com/shishao/hello/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = real.Retriever{}
	fmt.Println(download(r))
}