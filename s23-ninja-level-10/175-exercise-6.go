package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for j := 0; j < 10; j++ {
			c <- j
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
