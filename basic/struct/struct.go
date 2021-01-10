package main

import "fmt"

type Person struct {
	name string
	age int
	email string
}

func main() {
	// 初始化
	person := Person{"Tom", 31, "tom@gmail.com"}
	person = Person{name:"Tom", email:"Tom@gmail.com", age: 31}

	fmt.Println(person) // 输出 {Tom 31 tom@gmail.com}

	pPerson := &person
	fmt.Println(pPerson) // 输出 &{Tom 31 tom@gmail.com}

	pPerson.age = 35
	person.name = "Jerry"
	fmt.Println(person) // 输出 {Jerry 35 tom@gmail.com}
}
