package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	if x {
		fmt.Println("x is true")
	} else {
		fmt.Println("x is false")
	}

	a := 7
	b := 42
	fmt.Println(a == b)
	fmt.Printf("%T\n", a == b)
}
