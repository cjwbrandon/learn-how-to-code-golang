package main

import "fmt"

// Maps -> key-value pairs
// Note: Unordered

func main() {
	// map[(KEYTYPE)]{VALUETYPE}{{KEYS}: {VALUES}}
	m := map[string]int{
		"James": 13,
		"Bob":   20,
	}
	fmt.Println(m)
	fmt.Println(m["Bob"])

	// Entering a unregistered key -> zero value
	fmt.Println(m["abc"])

	if v, ok := m["abc"]; ok {
		fmt.Println("abc exists in m", v)
	}
}
