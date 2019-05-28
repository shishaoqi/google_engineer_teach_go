package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func printSliceValue(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	var s []int //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSliceValue(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSliceValue(s2)
	printSliceValue(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSliceValue(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSliceValue(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
	printSliceValue(s2)
}
