package main

//通过向channel发送值进行一对一通知

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)
// 如果一个channel没有可接收的值（即为空）,则该channel上的"接收操作"将阻塞,直到另一个goroutine将值发送到该channel.
// 因此,我们可以将值发送到channel,以通知另一个正在等待从同一channel接收值的 goroutine.
//
// 在以下示例中,该channel done 用作执行通知的signal channel.
func main() {
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}, 10) // 可以是缓存channel 也可以不是

	// goroutine 排序
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{} // ===== 发送操作
	}()

	// do some other things ...

	<- done // waiting here for notification ===== 接收操作
	fmt.Println(values[0], values[len(values)-1])
}
