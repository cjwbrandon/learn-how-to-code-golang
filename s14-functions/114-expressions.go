package main

import "fmt"

func main() {
	// Function Expressions
	// Functions are first-class citizens -> Can be used like any other variables/types
	f := func() {
		fmt.Println("my first func expressions")
	}
	f()

	f2 := func(x int) {
		fmt.Println(x)
	}
	f2(32142)
}
