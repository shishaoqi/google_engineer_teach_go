package main

import (
	"fmt"
	"reflect"
)

var globalVar = 3
var globalStr = "hello"
// bb := true // 此种是不行

var (
	a = 1
	b = "a"
	c = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func varibleShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s, globalStr)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	varibleShorter()

	str := fmt.Sprintf("%d - %s - %v", a, b, c)
	fmt.Println(str)
	fmt.Println("type:", reflect.TypeOf(str)) // data type check
}