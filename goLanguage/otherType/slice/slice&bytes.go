package main

import "fmt"

func main() {
	// 从字符串 获取 字节切片(字符串转换成字节数组)
	// 1. sl := []byte(s)
	// 2. copy函数：copy(dst byte[], src string)

	s := "\u00ff\u754c"

	sl := []byte(s)
	fmt.Println(sl)

	for i, c := range s {
		fmt.Printf("index: %d - character: %c\n", i, c)
	}
}