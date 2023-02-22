package main

import (
	"fmt"
	"shishaoGo/goLanguage/otherType/interface/mock"
	"shishaoGo/goLanguage/otherType/interface/real" // 注意点： // Retriever 是指针
	"time"
)

type Retriever_pointer interface {
	Get(url string) string
}

func download(r Retriever_pointer) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever_pointer // 接口类型变量

	//------------- 测试1 ----------------
	// 声明指针接口
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetrieverP := r.(*real.Retriever) // 通 .(*类型名字) 来取得interface肚子里的类型
	fmt.Println(realRetrieverP.TimeOut)

	// 类型可以通过 switch 来判断，也可以通过 type assertion 来判断
	// Type assertion 如果是 mock.Retriever
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock interface")
	}

	//------------- 测试2 ----------------
	// 声明（非指针）接口
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever) // 通 .(*类型名字) 来取得interface肚子里的类型
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock interface")
	}

	//fmt.Println(download(r))
}

// 接口的肚子里有两个东西：类型 与 值
// 类型可以通过 switch 来判断
func inspect(r Retriever_pointer) {
	fmt.Printf("%T ---- %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)

	case *real.Retriever:
		fmt.Println("Contents:", v.UserAgent)
	}
}
