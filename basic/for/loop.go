package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// for init; condition; post { }
// nit： 一般为赋值表达式，给控制变量赋初值；
// condition： 关系表达式或逻辑表达式，循环控制条件；
// post： 一般为赋值表达式，给控制变量增量或减量。
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// read file at line
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(convertToBin(5)) // 101

	fmt.Println(convertToBin(8)) // 1101

	fmt.Println(convertToBin(28))

	printFile("basic/abc.txt")
}
