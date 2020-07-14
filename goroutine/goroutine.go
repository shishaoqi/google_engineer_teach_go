package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		index := i
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i) // 如果这边直接引用外面的 i 不安全, 故加 index
			}
		}(index)
	}

	time.Sleep(time.Microsecond)
}
