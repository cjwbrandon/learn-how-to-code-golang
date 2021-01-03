package main

import "fmt"

func main() {
	// initialise
	c := make(chan int)

	// send
	go foo(c)

	// receive -> Waits/Blocks until it can do what it does (alternatively u can use wait groups)
	bar(c)
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
