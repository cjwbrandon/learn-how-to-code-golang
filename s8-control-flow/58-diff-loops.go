package main

import "fmt"

func main() {
	// For statements with single condition -> Similar to while
	fmt.Println("For Statement with single condition")
	x := 1
	for x < 5 {
		fmt.Println(x)
		x++
	}
	fmt.Println()

	// For statement with for clause
	fmt.Println("For Statement with for clause")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// For statements with range clause - future

	// For - infinite loop
	fmt.Println("For - Infinite loop")
	y := 1
	for {
		if y >= 5 {
			break // explicitly break
		}
		fmt.Println(y)
		y++
	}
}
