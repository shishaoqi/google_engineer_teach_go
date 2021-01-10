package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt2 struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt2) increment2() {
	/*a.lock.Lock()
	defer a.lock.Unlock()

	a.value++*/
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt2) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt2

	for i := 0; i < 100; i++ {
		go func() {
			a.increment2()
		}()
	}

	for i := 0; i < 100000; i++ {
		a.increment2()
	}

	go func() {
		for i := 0; i < 100; i++ {
			a.increment2()
		}
	}()

	for i := 0; i < 100000; i++ {
		a.increment2()
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
