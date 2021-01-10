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

	for i := 0; i < 100; i++ {
		go func() {
			a.increment()
		}()
	}

	for i := 0; i < 100000; i++ {
		a.increment()
	}

	go func() {
		for i := 0; i < 100; i++ {
			a.increment()
		}
	}()

	for i := 0; i < 100000; i++ {
		a.increment()
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
