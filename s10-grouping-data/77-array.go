package main

import "fmt"

func main() {
	// var {IDENTIFIER} [{SIZE}]{TYPE}
	var x [5]int
	fmt.Println(x)

	// 0-based index
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
