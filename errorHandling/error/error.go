package main

import (
	"errors"
	"fmt"
)

/*
通常情况下，错误的类型是error，它有一个内建的接口。

type error interface {
    Error() string
}
 */

// 自定义的出错结构
type myError struct {
	arg int
	errMsg string
}

// 实现 Error 接口
func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}
//两种出错
func error_test(arg int) (int, error) {
	if arg < 0  {
		return -1, errors.New("Bad Arguments - negtive!")
	}else if arg >256 {
		return -1, &myError{arg, "Bad Arguments - too large!"}
	}
	return arg*arg, nil
}
//相关的测试
func main() {
	for _, i := range []int{-1, 4, 1000} {
		if r, e := error_test(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("success:", r)
		}
	}
}