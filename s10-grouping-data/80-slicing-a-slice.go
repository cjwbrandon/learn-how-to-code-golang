package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x[1])
	fmt.Println(x[1:3])
	fmt.Println(x[2:7]) // Must be within length
}
