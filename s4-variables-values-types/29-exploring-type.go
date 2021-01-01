package main

import "fmt"

// STATIC programming Language
// A VARIABLE is DECLARED ot hold a VALUE of a certain TYPE
// DECLARE the VARIABLE with the IDENTIFIER {} of TYPE {} and ASSIGNED the VALUE {}
var y = 42
var z string = "abc"
var a string = `halo
"Using back-ticks"`

func main() {
	fmt.Println(y)
	// Print type
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	// Print type
	fmt.Printf("%T\n", z)

	// Error -> Cannot reassign to a different type
	// z was DECLARED with TYPE string
	// z = 43

	fmt.Println(a)
}
