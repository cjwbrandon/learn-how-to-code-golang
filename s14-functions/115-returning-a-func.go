package main

import "fmt"

func main() {
	f := foo()
	f("abc")

	f2 := bar()
	x := f2()
	fmt.Println(x)
}

// Returning a function
// Why? Callback functions -> Taking Functions as a Parameter
func foo() func(s string) {
	f := func(s string) {
		fmt.Println("Created by foo:", s)
	}
	return f
}

func bar() func() int {
	f := func() int {
		return 123
	}
	return f
}
