//大小写分开打
package main

import (
	"fmt"
	"sync"
)

type worker6 struct {
	in chan int
	done func()
}

func doWork6(id int, c chan int, w worker6) {
	for n := range c {
		fmt.Printf("Worker %d reveived %c\n", id, n)

		w.done()
	}
}

func createWorker6(chanId int, wg *sync.WaitGroup) worker6 {
	w := worker6{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork6(chanId, w.in, w)
	return w
}

func chanDemo6() {
	var workers [10]worker6
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker6(i, &wg)
	}

	wg.Add(20)
	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo6()
}

