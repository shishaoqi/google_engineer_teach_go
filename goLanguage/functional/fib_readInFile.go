package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

// 第一步: 创建类型 intGen (即一个函数的类型，但不管怎么样，是一个类型，它就能够实现接口。
// 这就 go 语言的灵活的地方)
type intGen func() int
// 函数它也能实现接口，那函数为什么能实现接口呢？go 语言的接收者 Receiver 它其实是个普通的函数参数，它没有很大的一个区别，
// 只是它写在了前面，语法上不一样

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	printFileContents(f)
}