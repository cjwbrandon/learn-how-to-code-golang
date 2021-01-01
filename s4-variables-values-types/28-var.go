package main

import "fmt"

// Error - bad
// x := 43

// Able - scope of y would be package level (global)
var y = 43

// DECLARE a VARIABLE with the IDENTIFIER "z"
// Of TYPE int
// ASSIGNS 0 of TYPE int to "z"
var z int // assigns 0 by default

func main() {
	x := 30
	fmt.Println(x)

	// The var keyword
	fmt.Println(y)

	// Best practice -> Limit the scope of your variables
	// Use short declaration operator as much as possible
	// y := 42

	fmt.Println(z)
}
