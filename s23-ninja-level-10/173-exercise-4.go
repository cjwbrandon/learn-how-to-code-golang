package main

import "fmt"

func main() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	receive(c, q)
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Channel:", v)
		case <-q:
			fmt.Println("quit")
			return
		}
	}
}
