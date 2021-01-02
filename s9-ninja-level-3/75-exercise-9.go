package main

import "fmt"

func main() {
	switch favSport := "abc"; favSport {
	case "abc":
		fmt.Println("asfdas")
	default:
		fmt.Println("default")
	}
}
