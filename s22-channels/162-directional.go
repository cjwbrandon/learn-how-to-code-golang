package main

import "fmt"

// Only can receive or only can send to this Channel
func main() {
	// Bi-directional
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)

	// Uni-directional
	// Send-only channel -> You can only send values to it
	d := make(chan<- int, 2)
	d <- 42
	d <- 43
	// fmt.Println(<-d) // Error
	fmt.Printf("%T\n", d)

	// Receive-only
	e := make(<-chan int, 2)
	// e <- 40 // Error
	fmt.Printf("%T\n", e)

	// Uni-directional -> Bi-directional works
	d = c
	e = c

	// Bi-directional -> Uni-directional Error
	// c = d
	// c = e
}
