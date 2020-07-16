package main

import (
	"bufio"
	"fmt"
	"github.com/shishao/hello/errorHandling/functional/fib"
	"os"
)

func writeFile_2(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println("file already exists")
		fmt.Println("Error:", err) // err.Error()
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// 把 非波那qi 写入文件（主要是学习 defer 的使用）
	writeFile_2("fib.txt")
}
