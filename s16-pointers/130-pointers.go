package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)              // Shows address in memory
	fmt.Printf("%T %T\n", a, &a) // * represent pointer

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // Shows value in address -> Dereference the address

	// 2 variables referencing the same address
	*b = 43
	fmt.Println(a) // Both values are changed
}
