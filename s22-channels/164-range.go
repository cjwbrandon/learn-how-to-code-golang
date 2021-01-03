package main

import "fmt"

func main() {
	// initialise
	c := make(chan int)

	// send
	go foo(c)

	// receive
	// range -> Continues looping a channel until it is closed
	for v := range c {
		fmt.Println(v)
	}
}

// send
func foo(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("t", i)
	}
	close(c) // close channel
}
