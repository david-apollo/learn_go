package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "ccmouse",
		"course": "golang",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	courseName := m["course"]
	fmt.Println(`m["course"]`, courseName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	delete(m, "name")
	name, ok := m["name"]
	fmt.Println(name, ok)
}