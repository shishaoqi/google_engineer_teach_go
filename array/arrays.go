package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3} //指定数量
	arr3 := [...]int{2, 4, 6, 8, 10} //不指定数量
	var grid [4][5]int //数量写在类型前面

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// range 关键字使用
	for i, val := range arr3 {
		//fmt.Println(arr3[i])
		fmt.Println(i, val)
	}
}
