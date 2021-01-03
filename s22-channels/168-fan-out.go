package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	// Throttling -> limiting number of go routines
	for i := 0; i < goroutines; i++ {
		go func() {
			// Fanning out
			for v := range c1 {
				func(v2 int) {
					// Fanning in
					c2 <- timeConsumingWork(v2)

				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
