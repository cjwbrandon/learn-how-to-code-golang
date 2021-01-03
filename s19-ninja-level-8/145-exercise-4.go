package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 1, 2, 5, 6, 7, 8, 4, 2}
	xs := []string{"b", "e", "m", "d", "k"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
