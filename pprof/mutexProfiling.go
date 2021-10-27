package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

// 1.
// 需要特别注意的是 runtime.SetMutexProfileFraction 语句，如果未来希望进行互斥锁的采集，那么需要通过调用该方法来设置采集频率，
// 若不设置或没有设置大于 0 的数值，默认是不进行采集的。

// 2.
// 与 Mutex 的 runtime.SetMutexProfileFraction 相似，Block 也需要调用 runtime.SetBlockProfileRate() 进行采集量的设置，
// 否则默认关闭，若设置的值小于等于 0 也会认为是关闭。
func init() {
	//runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}


/**
	什么情况下会造成阻塞呢，一般有如下方式：
		1. 调用 chan（通道）
		2. 调用 sync.Mutex （同步锁)
		3. 调用 time.Sleep() 等等。
 	那么为了验证互斥锁的竞争持有者的堆栈跟踪，我们可以根据以上的 sync.Mutex 方式，来调整先前的示例代码
 */
func main() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 1000999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}

	_ = http.ListenAndServe(":6060", nil)
}

/**
Mutex Profiling
-----------------------------------------
go tool pprof 进行分析
-----------------------------------------
go tool pprof http://localhost:6060/debug/pprof/mutex

# top
# list
在输出的分析中比较准确的看到引起互斥锁的函数在哪里，锁开销在哪里

*/

/**
Block Profiling
-----------------------------------------
go tool pprof 进行分析
-----------------------------------------
go tool pprof http://localhost:6060/debug/pprof/block

# top
# list
在输出的分析中比较准确的看到阻塞情况

*/