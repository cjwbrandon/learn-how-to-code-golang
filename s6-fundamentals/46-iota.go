package main

import "fmt"

// iota, increases by 1

const (
	a = iota + 1
	b = iota
	c = iota
)

// resets
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T\n", a, b, c)
	fmt.Println(d, e, f)
}
