package main

import (
	"fmt"
	"shishaoGo/goLanguage/basic/init/b_package"
	"shishaoGo/goLanguage/basic/init/c_package"
)

var T int64 = funA()

// 每个源文件都可以包含有且只包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行
func init() {
	fmt.Println("init in main.go")
}

func funA() int64 {
	fmt.Println("calling funA()")
	return 64
}

// 初始化顺序： 变量初始化 > init() > main()
func main() {
	fmt.Println("calling main")
	fmt.Println(b_package.B_global_param, c_package.C_global_param)
}
