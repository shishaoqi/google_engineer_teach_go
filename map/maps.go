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

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	causeName, ok := m["cause"]
	fmt.Printf("%q", causeName)

	courseName, ok = m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok = m["cause"]; ok {
		fmt.Printf("%q", causeName)
	} else {
		fmt.Printf("key does not exist %q", causeName)
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
