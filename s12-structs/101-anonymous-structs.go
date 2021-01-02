package main

import "fmt"

func main() {
	// Anonymous struct -> Does not have a name and defined directly
	// Note: Used only once
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
