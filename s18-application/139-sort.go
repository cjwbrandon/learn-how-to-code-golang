package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 6, 2, 1, 2, 7, 8, 3}
	sort.Ints(xi) // Mutates variable
	fmt.Println(xi)

	xs := []string{"a", "d", "c", "e", "b"}
	sort.Strings(xs)
	fmt.Println(xs)
}
