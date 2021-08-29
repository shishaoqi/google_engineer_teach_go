package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 内建变量类型

// 1. bool, string
// 2. (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
// 3. byte, rune(类同于 char)
// 4. float32, float64, complex64, complex128

// 复数
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	rt := cmplx.Pow(math.E, 1i * math.Pi) + 1
	fmt.Printf("%.3f\n", rt)
}

func main() {
	euler()
}