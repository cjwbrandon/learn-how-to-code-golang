package main

import "fmt"

func main() {
	x := []string{"a", "b", "c"}
	fmt.Println(x)
	y := []string{"d", "e", "f"}
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)
}
