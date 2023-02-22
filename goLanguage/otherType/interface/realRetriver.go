package main

import (
	"fmt"
	"shishaoGo/goLanguage/otherType/interface/real"
	"shishaoGo/goLanguage/otherType/interface/real_pointer"
	"time"
)

// 声明接口
type Retriever_real interface {
	Get(url string) string
}

func download_2(r Retriever_real) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	//
	var r Retriever_real // 接口类型变量

	// ---------- 非指针类型 ------------
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值

	// ---------- 指针类型 ------------
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T ---- %v\n", r, r)

	//fmt.Println(download_2(r))

	var rp Retriever_real
	// rp = real_pointer.Retriever{} 会报错，请尝试
	rp = &real_pointer.Retriever{ // 知道为什么一定要用接口指针吗  --  1. 首先，接口方法的实现是用指针类型实现的（其它没区别）
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T ---- %v\n", rp, rp)
}
