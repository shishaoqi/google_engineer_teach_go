package main

import (
	"fmt"
	"shishaoGo/goLanguage/otherType/interface/mock"
)

type Retriever_faker interface {
	Get(url string) string
}

// 接口由使用者定义-- 例如这边的 Retriever
func download(r Retriever_faker) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever_faker // 接口类型变量

	r = mock.Retriever{"value receiver. this is a fake imooc.com"}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值
	fmt.Println(download(r))

	r = &mock.Retriever{"pointer receiver. this is a fake imooc.com"}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值
	fmt.Println(download(r))
}
