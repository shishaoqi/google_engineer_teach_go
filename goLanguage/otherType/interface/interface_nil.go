/*
*

	 Equality
		Two interface values are equal

		1. if they have equal concrete values and identical dynamic types,
		2. or if both are nil.

		A value t of interface type T and a value x of non-interface type X are equal if

		1. t’s concrete value is equal to x
		2. and t’s dynamic type is identical to X.
*/
package main

import (
	"fmt"
)

type MyStringer interface {
	String() string
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main() {
	var x MyStringer
	fmt.Println(x == nil) // true
	fmt.Printf("%T ---- %v\n", x, x)

	// In the second print statement, the concrete value of x equals nil, but its dynamic type is *Point, which is not nil.
	x = (*Point)(nil)
	fmt.Println(x == nil) // false
	fmt.Printf("%T ---- %v\n", x, x)
}
