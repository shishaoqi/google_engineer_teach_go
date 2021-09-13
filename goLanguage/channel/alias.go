package main

import (
	"fmt"
	"time"
)

type chanAlias chan int

func main() {
	c := make(chanAlias)

	//c <- 1 // 如果这个在 go func(){} 会形成: fatal error: all goroutines are asleep - deadlock!
	go func() {
		fmt.Println(<- c)
	}()

	c <- 1
	time.Sleep(time.Second)
}