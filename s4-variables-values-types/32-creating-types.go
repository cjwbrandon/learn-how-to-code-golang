package main

import "fmt"

var a int

// Creating a type
// Creating TYPE {} of underlying TYPE {}
type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// Own type
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Note -> Error different types
	// a = b
}
