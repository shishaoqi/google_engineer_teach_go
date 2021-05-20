package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {
	Name string `json:"name"` //通过 struct tag 定义
	Age int	`json:"age"`
	Email string `json:"email"`
}

type PersonSlice struct {
	Persons []Person2
}

func main() {
	var ps PersonSlice

	ps.Persons = append(ps.Persons, Person2{"Tom", 31, "Tom@gmail.com"})
	ps.Persons =  append(ps.Persons, Person2{"Mike", 40, "Mike@gmail.com"})

	// 生成　ＪＳＯＮ
	json, err := json.Marshal(ps)
	if err != nil {
		fmt.Println("json err", err)
	}

	fmt.Println(string(json))
}
