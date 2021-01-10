package main

import "fmt"

type rect struct {
	width, height int
}

// Go语言中没有public, protected, private的关键字，
// 所以，如果你想让一个方法可以被别的包访问的话，你需要把这个方法的第一个字母大写。这是一种约定。
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2*(r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 15}
	fmt.Println("面积: ", r.area())
	fmt.Println("周长: ", r.perimeter())

	rp := &r
	fmt.Println("面积: ", rp.area())
	fmt.Println("周长: ", rp.perimeter())
}


