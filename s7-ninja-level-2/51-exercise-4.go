package main

import "fmt"

func main() {
	x := 2
	fmt.Printf("%d %b %x %#x\n", x, x, x, x)

	y := x << 1
	fmt.Printf("%d %b %x %#x\n", y, y, y, y)
}
