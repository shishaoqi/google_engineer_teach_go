package main

import "fmt"

func variableZeroValue() {
	var a int //变量名在前，变量类型在后
	var s string
	fmt.Printf("%d  %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4 // 声明初始化多个变量
	temp := "a string"
	var s string = "abc"
	fmt.Println(a, b, temp, s)
}

func main() {
	fmt.Println("Hello, world")
	variableZeroValue()
	variableInitialValue()
}
