package main

import "fmt"

func main() {
	f := even()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func even() func() int {
	x := 0
	return func() int {
		x += 2
		return x
	}
}
