package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < 10; j++ {
					c <- j
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
