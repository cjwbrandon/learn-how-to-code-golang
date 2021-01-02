package main

import "fmt"

func main() {
	// Set size of Slice
	// make([]{TYPE}, {length}, {capacity})
	x := make([]int, 3, 5)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[3] = 4 // Out of range
	x = append(x, 4)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 5)
	x = append(x, 6) // Reach capacity -> auto increase by 2x
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
