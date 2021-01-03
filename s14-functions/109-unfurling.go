package main

import "fmt"

// Unfurling a slice
func main() {
	xi := []int{1, 2, 3, 4, 5}
	x := sum(xi...)
	fmt.Println(x)

	fmt.Println(sum())
}

// Note: Variadic -> 0 or more
func sum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("Index: %d, Add: %d, Total: %d\n", i, v, sum)
	}

	return sum
}
