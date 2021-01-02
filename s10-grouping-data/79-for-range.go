package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	// for {index}, {value} := range {IDENTIFIER}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
