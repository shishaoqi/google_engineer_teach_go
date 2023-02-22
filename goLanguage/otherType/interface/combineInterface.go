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
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T ---- %v\n", r, r) // 接口的肚子里有两个东西：类型 与 值
	fmt.Println(download(r))

	var p Poster // 接口类型变量
	p = post.Post{"http://hello", make(map[string]string)}
	fmt.Println(postFun(p))

	retriever := &mock.Retriever{"this is a new fake imooc.com"}
	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}
