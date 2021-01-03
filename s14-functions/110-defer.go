package main

import "fmt"

func main() {
	foo()
	bar()

	// defer -> ran after the surrounding function returns, i.e. main()
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
