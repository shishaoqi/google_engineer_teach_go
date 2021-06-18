package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}

	return g
}

func main() {
	const filename = "basic/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
			grade(0),
			grade(59),
			grade(60),
			grade(82),
			grade(99),
			grade(100),
			grade(100),
		)
}