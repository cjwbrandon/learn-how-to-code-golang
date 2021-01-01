package main // must have

import "fmt"

// Entry point to your program
func main() {
	fmt.Println("Hello World")
	foo()

	// Iterative
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		// Conditional
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

// Control flow: (How a computer read/execute codes)
// (1) Sequence
// (2) Loop: Iterative
// (3) Conditional
