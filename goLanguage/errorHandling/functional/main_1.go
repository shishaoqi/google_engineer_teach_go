package main

import (
	"bufio"
	"fmt"
	"shishaoGo/goLanguage/errorHandling/functional/fib"
	"io"
	"os"
	"strings"
)

type intFun func() int

func (i intFun) Read(p []byte) (n int, err error) {
	next := i()
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents_2(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for i := 0; i < 15 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
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
	// 把 非波那qi 当文件来读
	var f intFun = fib.Fibonacci()
	printFileContents_2(f)

	// 把 非波那qi 写入文件（主要是学习 defer 的使用）
	writeFile("fib.txt")
}
