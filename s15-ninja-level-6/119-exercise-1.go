package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 123
}

func bar() (int, string) {
	return 456, "abc"
}
