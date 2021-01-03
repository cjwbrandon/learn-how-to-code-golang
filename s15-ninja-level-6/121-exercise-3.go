package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Run this at the end")
	}()
	func() {
		fmt.Println("Work work...")
	}()
	func() {
		fmt.Println("Ending work")
	}()
}
