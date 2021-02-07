package main

import (
	"fmt"
	"regexp"
)

func main()  {
	addr := "建省厦门市思明区厦门同步网络有限公司(软件园二期望海路10号之三3F)"

	reg := regexp.MustCompile(`\(.*?\)`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	result := reg.FindAllStringSubmatch(addr, -1)
	fmt.Println("result = ", len(result))
}


