package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
