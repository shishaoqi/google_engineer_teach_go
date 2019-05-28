package main

import (
	"fmt"
	"googlelearngo/real"
	"googlelearngo/mock"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name": "ccomuse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked imocc.com",
	})

	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	inspect(&retriever)

	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}
	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)

	fmt.Printf("> %T %v\n", r, r)
	fmt.Printf("> Type switch:")

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)

	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
