package main

import "fmt"

func main() {
	// Short Declaration Operator - :=
	// Declare variable x and assign a value/expression, 42
	x := 42
	fmt.Println(x)

	// Already been declared -> reassigning -> =
	x = 99

	y := 100 + 9
	fmt.Println(y)
}
