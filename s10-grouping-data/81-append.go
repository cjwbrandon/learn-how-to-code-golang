package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	// Append returns a new slice
	// Appends to end
	x = append(x, 6)       // append single
	x = append(x, 7, 8, 9) // append multiple
	fmt.Println(x)

	y := []int{10, 11, 12}
	// {identifier}...
	z := append(x, y...)
	fmt.Println(z)
}
