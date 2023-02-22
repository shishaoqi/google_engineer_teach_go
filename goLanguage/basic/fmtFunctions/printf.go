package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%t\n", 1 == 2) // 布尔占位符
	fmt.Printf("二进制：%b\n", 255)
	fmt.Printf("八进制：%o\n", 255)
	fmt.Printf("十进制：%d\n", 255)
	fmt.Printf("十六进制：%X\n", 255)
	fmt.Printf("浮点数：%f\n", math.Pi)
	fmt.Printf("字符串：%s\n", "hello world")

	type Human struct {
		Name string
	}

	var people = Human{Name: "shishao"}

	// 普通占位符
	fmt.Printf("相应值的默认格式 %v\n", people)
	fmt.Printf("打印结构体时，会添加字段名 %+v\n", people)
	fmt.Printf("相应值的Go语法表示 %#v\n", people)
	fmt.Printf("相应值的类型的Go语法表示 %T\n", people)
	fmt.Printf("字面上的百分号，并非值的占位符 %%\n")

	// 布尔占位符
	fmt.Printf("布尔占位符 %t\n", true)

	// 整数占位符
	fmt.Printf("二进制表示 %b\n", 5)
	fmt.Printf("相应Unicode码点所表示的字符 %c\n", 0x4E2D)
	fmt.Printf("八进制表示 %o\n", 10)
	fmt.Printf("十进制表示 %d\n", 0x12)
	fmt.Printf("十六进制表示，字母形式为小写 a-f %x\n", 13)
	fmt.Printf("十六进制表示，字母形式为大写 A-F %X\n", 13)
	fmt.Printf("Unicode格式: U+1234 %U\n", 0x4E2D)
	fmt.Printf("单引号围绕的字符字面值, 由Go语法安全地转义 %q\n", 0x4E2D)

	// 指针
	fmt.Printf("十六进制表示，前缀 0x %p", &people)
}
