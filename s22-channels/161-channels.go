package main

import "fmt"

// Bi-directional Channels
func main() {
	// chan for chnanel to put int
	c := make(chan int)

	// Channels block
	// c <- 42
	// Solution 1
	go func() {
		// Putting on a channel
		c <- 42
	}()

	// Taking off the channel
	fmt.Println(<-c)

	// Solution 2, buffer: make(chan, bufferChann)
	// try to avoid buffer due to complicated design
	d := make(chan int, 1)
	d <- 43
	fmt.Println(<-d)

	// Unsuccessful buffer
	// e := make(chan int, 1)
	// e <- 43
	// e <- 44 // Not enough space
	// fmt.Println(<-e)
	e := make(chan int, 1)
	e <- 43
	fmt.Println(<-e)
	e <- 44
	fmt.Println(<-e)
}
