package main

import (
	"log"
	"net/http"
	"time"
	_ "net/http/pprof"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d",  Add("go toure"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}

/**
		交互式终端下使用 Pprof
-----------------------------------------
flat：	函数自身的运行耗时。
flat%：	函数自身在 CPU 运行耗时总比例。
sum%：	函数自身累积使用 CPU 总比例。
cum：	函数自身及其调用函数的运行总耗时。
cum%：	函数自身及其调用函数的运行耗时总比例。
Name：	函数名。
-----------------------------------------

# CPU Profiling
go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60

1. top10
2. web  --  保存文件： CPU-Profiling-pprof001.svg
======================================================================================================


# Heap Profiling
go tool pprof http://localhost:6060/debug/pprof/heap

需要注意的一点是 Type 这一个选项，你可以看到它默认显示的是 inuse_space，实际上可以针对多种内存概况进行分析，常用的类别如下：
go tool pprof -inuse_space -- inuse_space：分析应用程序的常驻内存占用情况。
go tool pprof -alloc_objects --  alloc_objects：分析应用程序的内存临时分配情况

-- traces命令
======================================================================================================


# Goroutine Profiling
go tool pprof http://localhost:6060/debug/pprof/goroutine

-- traces命令
在查看 goroutine 时，我们可以使用traces命令，这个命令会打印出对应的所有调用栈，以及指标信息，可以让我们很便捷的查看到整个调用链路有什么，
分别在哪里使用了多少个 goroutine，并且能够通过分析查看到谁才是真正的调用方 ---- 在调用栈上来讲，其展示顺序是自下而上的
======================================================================================================


 */
