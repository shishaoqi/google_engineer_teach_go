package main

//通过向channel发送值进行一对一通知

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 可以是缓存channel 也可以不是

	// goroutine 排序
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
	}()

	// do some other things ...

	<- done // waiting here for notification =====
	fmt.Println(values[0], values[len(values)-1])
}
