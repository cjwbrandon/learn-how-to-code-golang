package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// Race condition: Each goroutine accessing a shared variable
	counter := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			v := counter
			runtime.Gosched() // Yield thread for other goroutines
			v++
			counter = v
			fmt.Println(runtime.NumGoroutine())
			wg.Done()
			fmt.Println(runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
