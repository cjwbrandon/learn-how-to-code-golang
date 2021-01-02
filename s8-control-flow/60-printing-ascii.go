package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("binary: %d\thexdecimal: %#x\tASCII Character: %#U\n", i, i, i)
	}
}
