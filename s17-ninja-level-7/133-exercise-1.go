package main

import "fmt"

var x int

func main() {
	fmt.Println(x, &x)
	x = 2
	fmt.Println(x, &x)
}
