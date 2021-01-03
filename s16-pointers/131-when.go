package main

import "fmt"

func main() {
	x := 0
	foo(x)
	fmt.Println("Value is not mutated", x)

	fmt.Println("Pointer")
	bar(&x)
	fmt.Println("Value of x is mutated:", x)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(y, *y)
	*y = 43
	fmt.Println(y, *y)
}
