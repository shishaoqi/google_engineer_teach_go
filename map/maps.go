package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "shishao",
		"course": "golang",
		"site": "shishao.com",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int // m3 == nil
	fmt.Println(m2, m3)

	fmt.Println("Travesing map")
	for k, v  := range m {
		fmt.Println(k, v)
	}

	fmt.Println("------- Getting values -------")
	courseName := m["course"]
	fmt.Println(courseName)
	causeName, ok := m["cause"]
	fmt.Println(causeName, ok)

	courseName, ok = m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok = m["cause"]; ok {
		fmt.Println(causeName)
		fmt.Println("ok: ", ok)
	} else {
		fmt.Println("cause key does not exist")
		fmt.Println("ok: ", ok)
	}

	fmt.Println("------- Deleting values -------")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("%q, %v", name, ok)
}
