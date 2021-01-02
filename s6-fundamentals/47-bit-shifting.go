package main

import "fmt"

// Using iota
const (
	_   = iota
	kb2 = 1 << (iota * 10)
	mb2 = 1 << (iota * 10)
	gb2 = 1 << (iota * 10)
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	// Bit shift by 1 -> Creates 4
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb) // Shifts 10 from kb
	fmt.Printf("%d\t%b\n", gb, gb)   // Shifts 10 from mb

	fmt.Printf("%d\t\t%b\n", kb2, kb2)
	fmt.Printf("%d\t\t%b\n", mb2, mb2) // Shifts 10 from kb
	fmt.Printf("%d\t%b\n", gb2, gb2)   // Shifts 10 from mb
}
