package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// 目标字符串
	searchIn := "John: 2578.34 Willian: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"

	f := func(s []byte) []byte {
		float, _ := strconv.ParseFloat(string(s), 32)
		return []byte(strconv.FormatFloat(float * 2, 'f', 2, 32))
	}

	re, _ := regexp.Compile(pattern)
	str := re.ReplaceAllFunc([]byte(searchIn), f)
	fmt.Println(string(str))
}
