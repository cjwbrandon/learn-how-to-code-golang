package main

import "fmt"

// Slice - Composite Literal
func main() {
	// Composite Literal -> {IDENTIFIER} := {TYPE}{{VALUES}}
	// Note: a SLICE allows you to group together VALUES of the same TYPE
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
}
