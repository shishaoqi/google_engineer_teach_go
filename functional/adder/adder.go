package main

import "fmt"

// 闭包
// 函数体包函 局部变量（参数也是局部变量）

// sum 是这个函数所处的环境里的变量，叫 自由变量，会保存它的状态 ----编义器会连一根线
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// "正统"函数式编程
// 1. 不可变性：不能有状态，只有常量和函数
// 2. 函数只能有一个参数
// 3. 本课不作上述严格规定

// 这是一个稍微正统一点的函数式写法
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}


func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
