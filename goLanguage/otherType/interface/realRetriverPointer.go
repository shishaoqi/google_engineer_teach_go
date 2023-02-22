/**
 * 此文件与 realRetriver.go 文件是一样的测试效果,
 * 只是声明了一个“不同的接口”
 */
package main

import (
	"fmt"
	"shishaoGo/goLanguage/otherType/interface/mock"
	"shishaoGo/goLanguage/otherType/interface/real"         // 注意点： 此结构体是按结构本身“接收者”实现接口方法的
	"shishaoGo/goLanguage/otherType/interface/real_pointer" // 此结构体是按结构的“指针接收者”实现接口方法的
	"time"
)

type Retriever_2 interface {
	Get(url string) string
}

func download(r Retriever_2) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever_2 // 接口类型变量

	//------------- 测试1 ----------------
	/* r = real_pointer.Retriever, 使用
	 * 报：cannot use (real_pointer.Retriever literal) (value of type real_pointer.Retriever) as Retriever_2 value
	 * in assignment: real_pointer.Retriever does not implement Retriever_2 (method Get has pointer receiver)compiler
	 * -- （InvalidIfaceAssign 错误类型）：解决办法是：给 real_pointer 加方法》》》 func (r Retriever) Get(url string) string {}
	 */
	r = &real_pointer.Retriever{
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
func inspect(r Retriever_2) {
	fmt.Printf("%T ---- %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)

	case *real.Retriever:
		fmt.Println("Contents:", v.UserAgent)
	}
}
