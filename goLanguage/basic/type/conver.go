package main

import (
	"fmt"
	"reflect"
)

func main() {
	var p interface{}

	p = 21

	fmt.Println(reflect.TypeOf(p))
}
