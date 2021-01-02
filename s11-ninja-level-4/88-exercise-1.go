package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Printf("%T\n", x)
}
