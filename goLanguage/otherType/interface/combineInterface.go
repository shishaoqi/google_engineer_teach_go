/**
 *	接口组合示例
 */
package main

import (
	"fmt"
	"shishaoGo/goLanguage/otherType/interface/mock"
	"shishaoGo/goLanguage/otherType/interface/post"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, from map[string]string) string
}

const url = "http://www.imooc.com"

// 一函数：使用接口类型的变量作为函数参数
func download(r Retriever) string {
	return r.Get(url)
}

func postFun(p Poster) string {
	return p.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

// 组合接口类型
type RetrieverPoster interface {
	Retriever
	Poster
	//Connect(host string)
}

func session(s RetrieverPoster) string {
	// s.Get("http:://imooc.com")
	// s.Post("http:://www.baidu.com", map[string]string {"name":"shishao", "addr":"xiamenng"})
	s.Post(url, map[string]string{
		"contents": "another facked imooc.com",
	})

	return s.Get(url)
}

func main() {
	var r Retriever // 接口类型变量

	r = &mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值
	fmt.Println(download(r))
	fmt.Println(session(r)) // 为什么这句不行？1. 如果把 mock.Retriever 实现接口方法 从 “指针接收者” 改为 “值接收者” 会如何？ 2. 不按 1 方法，还能如何办，让此句正常编译？
	// 答： 1 方案也不行；  没找到解决方案（只能：声明结构指针）

	var p Poster // 接口类型变量
	p = post.Post{"http://hello", make(map[string]string)}
	fmt.Println(postFun(p))

	retriever := &mock.Retriever{"this is a new fake imooc.com"}
	fmt.Printf("%T ---- %v\n", retriever, retriever)
	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}
