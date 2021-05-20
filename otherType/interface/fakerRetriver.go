package main

import (
	"fmt"
	"github.com/shishao/hello/otherType/interface/mock"
)

type Retriever_1 interface {
	Get(url string) string
}

// 接口由使用者定义-- 例如这边的 Retriever
func download(r Retriever_1) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever_1
	r = &mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值
	fmt.Println(download(r))
}