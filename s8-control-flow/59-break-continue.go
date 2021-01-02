package main

import "fmt"

// break - exit loop
// continue - stop and continue to the next iteration

func main() {
	x := 0
	for x < 5 {
		x++
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
}
