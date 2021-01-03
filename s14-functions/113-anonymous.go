package main

import "fmt"

func main() {
	// Anonymous Functions -> IIFE
	func() {
		fmt.Println("Anonymous func ran")
	}()
	// With parameters
	func(x int) {
		fmt.Println(x)
	}(42)
}
