package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func euler() {
	rt := cmplx.Pow(math.E, 1i * math.Pi) + 1
	fmt.Printf("%3f",rt)
}

func main() {
	euler()
}