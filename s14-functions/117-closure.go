package main

import "fmt"

// Package scoped
var x int = 0

func main() {
	fmt.Println(x)

	// Function scoped
	y := 1
	fmt.Println(y)

	// Closure using {} - limiting scope of z within function
	{
		z := 2
		fmt.Println(z)

		// Can assign used variables in different scopes
		y := 3
		fmt.Println(y) // Uses y within closure
	}
	fmt.Println(y) // Uses y within function scope

	fmt.Println()
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b()) // different scope
	fmt.Println(b())
}

// Another use case for closure
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
