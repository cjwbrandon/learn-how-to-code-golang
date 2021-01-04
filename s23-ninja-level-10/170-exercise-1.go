package main

import "fmt"

func main() {
	// Using goroutines
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// Using buffered channels
	d := make(chan int, 1)

	d <- 42

	fmt.Println(<-d)
}
