package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	// Deleting using append
	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}
