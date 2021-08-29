package main

import "fmt"

// 常量
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
}

func enums() {
	const(
		cpp 	= iota
		_
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
