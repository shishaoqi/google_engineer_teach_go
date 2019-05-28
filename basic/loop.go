package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

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
	fmt.Println(convertToBin(5))

	fmt.Println(convertToBin(8))

	fmt.Println(convertToBin(28))

	printFile("abc.txt")
}
