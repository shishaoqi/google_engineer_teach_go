package main

import "fmt"

type Fn func(int) bool

func main() {
	var slices  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := filter(slices, even)
	fmt.Println("even slice: ", s)
}

func filter(s []int, fn Fn) []int{
	evenSlice := []int{}

	for num := range s {
		if fn(num) {
			evenSlice = append(evenSlice, num)
		}
	}

	return evenSlice
}

func even(num int) bool {
	if num % 2 == 0 {
		return true
	}

	return false
}


