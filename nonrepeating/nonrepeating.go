package main

import (
	"fmt"
)


func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)

	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))

	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))

	fmt.Println(lengthOfNonRepeatingSubStr("黑灰化肥灰会挥发发灰黑讳为黑灰花会飞"))
}