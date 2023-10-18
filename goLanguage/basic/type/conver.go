package main

import (
	"fmt"
	"reflect"
	"sync"
)

type atomicInt2 struct {
	value int
	lock sync.Mutex
}

func main() {
	var p interface{}

	p = 21

	typeOfp := reflect.TypeOf(p)
	fmt.Println(typeOfp)

	fmt.Println(typeOfp.Name(), typeOfp.Kind())

	var a atomicInt2
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
