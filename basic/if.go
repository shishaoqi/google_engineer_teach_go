package main

import "fmt"

// if 的条件里不需要括号
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func main() {
	fmt.Println(bounded(1))

	fmt.Println(bounded(101))
}