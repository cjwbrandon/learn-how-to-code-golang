package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5)
	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)
}

// Variadic Params -> Unlimited number of param
// Note: Automatically gets converted to a slice
func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("Index: %d, Add: %d, Total: %d\n", i, v, sum)
	}

	return sum
}
