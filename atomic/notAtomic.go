package main

import (
	"fmt"
	"time"
)

type notAtomicInt int

func (a *notAtomicInt) increment() {
	*a++
}

func (a *notAtomicInt) get() int {
	return int(*a)
}

//go run -race notAtomic.go
func main() {
	var a notAtomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
