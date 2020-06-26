package main

import "fmt"
// go 数组是按值传递的
// 要实现改变的话，必须使用指针
// func printArray(arr *[5]int) {

func printArray(arr [5]int) {
	arr[0] = 100 // 想改变的地方
	for i, v := range arr {
		println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3} //指定数量
	arr3 := [...]int{2, 4, 6, 8, 10} //不指定数量
	var grid [4][5]int //数量写在类型前面

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	// range 关键字使用
	/*for i, val := range arr3 {
		//fmt.Println(arr3[i])
		fmt.Println(i, val)
	}*/

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
