package main

import (
	"fmt"
	"time"
)

func main() {
	// 2006-01-02       2006-01-02 15:04:05
	lastDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println(lastDay)
}
