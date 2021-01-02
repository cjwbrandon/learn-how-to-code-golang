package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 13,
		"Bob":   20,
	}
	fmt.Println(m)

	// Adding element
	m["Bobby"] = 21
	fmt.Println(m)

	// Looping using range
	for k, v := range m {
		fmt.Println(k, v)
	}
}
