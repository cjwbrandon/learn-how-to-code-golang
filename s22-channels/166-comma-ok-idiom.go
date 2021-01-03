package main

import "fmt"

// Select statement
func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		// comma ok idioms
		// Note: ok returns false when channel is empty or closed
		select {
		case v, ok := <-e:
			fmt.Println("Even Channel:", v, ok)
		case v, ok := <-o:
			fmt.Println("Odd Channel:", v, ok)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			}
			fmt.Println("from comma ok", i)
		}
	}
}
