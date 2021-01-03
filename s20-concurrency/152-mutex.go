package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// Race condition: Each goroutine accessing a shared variable
	counter := 0
	var wg sync.WaitGroup
	// Solution: Mutex
	var mu sync.Mutex

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock() // Lock for reading and writing
			v := counter
			runtime.Gosched() // Yield thread for other goroutines
			v++
			counter = v
			mu.Unlock() // Unlock once done
			fmt.Println(runtime.NumGoroutine())
			wg.Done()
			fmt.Println(runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
