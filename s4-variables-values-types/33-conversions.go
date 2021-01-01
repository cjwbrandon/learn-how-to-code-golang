package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// Own type
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Conversion not casting
	a = int(b)
	fmt.Println("New value of a: ", a)
	fmt.Printf("%T\n", a)
}
