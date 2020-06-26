package main

import "fmt"

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}

	return result
}

func main() {
	test_1 := eval(32, 8, "*")
	test_2 := eval(90, 10, "/")

	fmt.Printf("test_1=%v, test_2=%v", test_1, test_2)
}
