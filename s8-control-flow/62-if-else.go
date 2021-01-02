package main

import "fmt"

func main() {
	// You can stack as many else if as you want
	if x := 42; x == 40 {
		fmt.Println("This won't run")
	} else if x == 41 {
		fmt.Println("This is still false")
	} else {
		fmt.Println("Since both if statements are false, we run this")
	}
}
