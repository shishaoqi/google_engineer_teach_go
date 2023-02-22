package main

import (
	"fmt"
)

func main() {

	f := Add()
	fmt.Println(f())
	fmt.Println(f())

	f2 := Add()
	fmt.Println(f2())

	fmt.Println(f())
	fmt.Println(f())
}

func Add() func() int {
	i := 0 // 此函数内的局部变量始终保存在内存中，即只会声明闭包时调用一次

	return func() int {
		for n := 0; n < 3; n++ {
			i += 1
		}
		return i
	}
}
