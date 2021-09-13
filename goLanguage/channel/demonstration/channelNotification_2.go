package main

//通过从channel接收值进行一对一通知

import (
	"fmt"
	"time"
)

// 如果channel的值缓冲区队列已满（未缓冲channel的缓冲区队列始终已满）,则该channel上的"发送操作"将阻塞,
// 直到另一个goroutine从该channel接收到值为止. 因此,我们可以从一个channel接收一个值,
// 以通知另一个正在等待将值发送到同一channel的goroutine. 通常,该channel应为无缓冲channel.
//
// 与上一个示例中介绍的方式相比,使用这种通知方式的方式要少得多.
func main() {
	done := make(chan struct{})
	// The capacity of the signal channel can
	// also be one. If this is true, then a
	// value must be sent to the channel before
	// creating the following goroutine.

	go func() {
		fmt.Print("Hello")
		// Simulate a workload.
		time.Sleep(time.Second * 2)

		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		<- done
	}()

	// Blocked here, wait for a notification. === send opration is blocked
	done <- struct{}{}
	fmt.Println(" world!")
}
