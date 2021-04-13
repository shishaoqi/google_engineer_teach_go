package main

import (
	"fmt"
	"github.com/benmanns/goworker"
	"runtime"
)

func myFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args)
	return nil
}

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    100,
		Queues:         []string{"myqueue", "delimited", "queues"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "resque:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("MyClass", myFunc)
}

// redis 中执行： redis-cli -r 100 RPUSH resque:queue:myqueue '{"class":"MyClass","args":["hi","there"]}'
// 会有效果
func main() {
	defer func() {
		if r := recover(); r != nil {
			//打印调用栈信息
			buf := make([]byte, 2048)
			n := runtime.Stack(buf, false)
			stackInfo := fmt.Sprintf("%s", buf[:n])
			fmt.Printf("panic stack info %s", stackInfo)
		}
	}()

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}