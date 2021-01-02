package main

import "fmt"

func main() {
	// Pre-declared constant - bool
	if true {
		fmt.Println("IT IS TRUE")
	}

	if false {
		fmt.Println("IT IS FALSE")
	}
	if !true {
		fmt.Println("IT IS NOT TRUE")
	}

	if !false {
		fmt.Println("IT IS NOT FALSE")
	}

	// Using expressions
	if 2 == 2 {
		fmt.Println("2 is equal to 2")
	}

	if 2 != 3 {
		fmt.Println("Yes, 2 is not equal to 3")
	}

	// if init - useful for limiting scope of x
	fmt.Println("if init; condition")
	if x := 42; x == 42 {
		fmt.Println("Trying with init")
	}
}
