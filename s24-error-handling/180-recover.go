package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.") // Continues running rest in goroutine
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // Does not run this after defer
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i) // Last in First Out
	fmt.Println("Printing in g", i)
	g(i + 1)
}
