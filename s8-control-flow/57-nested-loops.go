package main

import "fmt"

func main() {
	// Outer loop
	for i := 0; i < 10; i++ {
		fmt.Printf("The Outer loop: %d\n", i)
		// Inner loop
		for j := 0; j < 3; j++ {
			fmt.Printf("\tThe inner loop: %d\n", j)
		}
	}
}
