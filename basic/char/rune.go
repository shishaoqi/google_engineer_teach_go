package main

import "fmt"

func main() {
	str := "hello 世界"

	fmt.Println("str string len:", len(str))
	fmt.Println("str rune len:", len([]rune(str)))
}
