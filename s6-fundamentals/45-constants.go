package main

import "fmt"

// const a = 42
// const b = 42.78
// const c = "James Bond"

// Declaring all the constatnts tgt
const (
	// Untyped constants
	a = 42
	b = 42.78
	c = "James Bond"

	// Typed constants
	// a int     = 42
	// b float64 = 42.78
	// c string  = "James Bond"
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
