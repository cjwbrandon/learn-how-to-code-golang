package main

import "fmt"

var g = func() {
	fmt.Println("Hello World from g")
}

func main() {
	f := func() {
		fmt.Println("Hello World")
	}
	f()
	g()
}
