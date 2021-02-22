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
	//err = errors.New("this is a custom error") // 创造出一个非 os.PathError

	if err != nil {
		//panic(err)
		//fmt.Println("file already exists")
		//fmt.Println("Error:", err) // err.Error()

		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s -- %s -- %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
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
