//大小写分开打
package main

import (
	"fmt"
	"sync"
)

type worker5 struct {
	in chan int
	wg *sync.WaitGroup // 改用 waitgroup
}

func doWork5(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d reveived %c\n", id, n)

		wg.Done()
	}
}

func createWorker5(chanId int, wg *sync.WaitGroup) worker5 {
	w := worker5{
		in: make(chan int),
		wg: wg,
	}
	go doWork5(chanId, w.in, wg)
	return w
}

func chanDemo5() {
	var workers [10]worker5
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker5(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo5()
}

