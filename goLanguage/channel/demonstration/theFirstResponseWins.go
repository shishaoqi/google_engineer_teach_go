package main


import (
	"fmt"
	"time"
	"math/rand"
)

func source(c chan<- int32, goroutinIndex int) {
	ra, rb := rand.Int31(), rand.Intn(3) + 1
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
	fmt.Println(goroutinIndex)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	// c must be a buffered channel.
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c, i)
	}
	// Only the first response will be used.
	rnd := <- c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}