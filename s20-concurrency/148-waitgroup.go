package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// main runs on 1 goroutine itself
func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	// Concurrent Code: Launching in its own goroutine
	// Parallelism possibility depends on NumCPU
	// Note: Diff go routine -> Cannot see output
	// Synchronization
	wg.Add(2) // Wait for 1 thing
	go foo()
	go bar()

	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	// Why? When the main exits, all goroutines will be closed regardless if they have exitted
	wg.Wait() // Wait
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // Thing is Done
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
