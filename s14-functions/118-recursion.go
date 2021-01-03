package main

import "fmt"

// Recursion -> Functions calling itself
// Must have some way to stop
// Recursion can usually be implemented using loops (clearer)
func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	// switch n {
	// case 1:
	// 	return 1
	// default:
	// 	return n * factorial(n-1)
	// }
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
