package main

import "fmt"

// Looks like Classes in JS
// But in Go, they are types
type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
