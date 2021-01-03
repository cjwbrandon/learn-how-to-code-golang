package main

import "fmt"

// Select statement
func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		// Using select to pull values off different channels
		select {
		case v := <-e:
			fmt.Println("Even Channel:", v)
		case v := <-o:
			fmt.Println("Odd Channel:", v)
		case v := <-q:
			fmt.Println("Quit Channel:", v)
			return
		}
	}
}
