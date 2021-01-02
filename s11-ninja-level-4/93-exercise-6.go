package main

import "fmt"

func main() {
	x := make([]string, 50, 50)
	x = []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(x); i++ {
		fmt.Printf("Index: %d, Country: %v\n", i, x[i])
	}
}
