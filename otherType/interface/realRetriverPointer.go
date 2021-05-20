package main

import (
	"fmt"
	"github.com/shishao/hello/otherType/interface/mock"
	"github.com/shishao/hello/otherType/interface/real_pointer"
	"time"
)

type Retriever_3 interface {
	Get(url string) string
}

func download_3(r Retriever_3) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever_3
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
		fmt.Println("not a mock interface")
	}

	//fmt.Println(download(r))
}

// 接口的肚子里有两个东西：类型 与 值
// 类型可以通过 switch 来判断
func inspect(r Retriever_3) {
	fmt.Printf("%T ---- %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)

	case *real.Retriever:
		fmt.Println("Contents:", v.UserAgent)
	}
}