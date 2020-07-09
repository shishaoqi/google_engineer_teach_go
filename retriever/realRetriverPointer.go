package main

import (
	"fmt"
	"github.com/shishao/hello/retriever/mock"
	"github.com/shishao/hello/retriever/real_pointer"
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
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever) // 通 .(*类型名字) 来取得interface肚子里的类型
	fmt.Println(realRetriever.TimeOut)

	// 类型可以通过 switch 来判断，也可以通过 type assertion 来判断
	// Type assertion 如果是 mock.Retriever
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}

// 接口的肚子里有两个东西：类型 与 值
// 类型可以通过 switch 来判断
func inspect(r Retriever) {
	fmt.Printf("%T ---- %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)

	case *real.Retriever:
		fmt.Println("Contents:", v.UserAgent)
	}
}