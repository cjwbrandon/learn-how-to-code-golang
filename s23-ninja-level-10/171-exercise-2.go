package main

import "fmt"

func main() {
	// Send-only channel
	cs := make(chan<- int)
	// Receive-only channel
	cr := make(<-chan int)

	fmt.Printf("%T %T", cs, cr)
}
