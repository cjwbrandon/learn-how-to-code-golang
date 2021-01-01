package main

import "fmt"

type int2 int

var x int2
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	// Conversion
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
