package main

import (
	"fmt"
)

func main() {
	slFrom := []int{1, 2, 3}
	slTo  := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n)
	fmt.Printf("Copied slice capasity: %d", cap(slTo))
}
