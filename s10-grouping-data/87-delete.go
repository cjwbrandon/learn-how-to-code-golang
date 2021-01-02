package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)

	delete(m, "b")
	fmt.Println(m)

	// Deleting a non-existing key -> no errors
	delete(m, "c")
	fmt.Println(m)

	if _, ok := m["c"]; ok {
		delete(m, "c")
	}
}
