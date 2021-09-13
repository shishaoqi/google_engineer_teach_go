package main

import (
	"fmt"
	"shishaoGo/goLanguage/otherType/interface/real"
	"time"
)

type Retriever_2 interface {
	Get(url string) string
}

func download_2(r Retriever_2) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever_2
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T ---- %v\n", r, r)

	fmt.Println(download_2(r))
}