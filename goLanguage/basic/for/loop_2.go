package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// read file at line
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// file 为什么能传给 io.Reader , 因为 *File 实现了 io.Reader 的接口
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	printFile("basic/abc.txt")
}
