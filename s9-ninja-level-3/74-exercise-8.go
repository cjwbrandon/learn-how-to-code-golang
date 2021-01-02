package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("NAHH")
	case ("a" == "a"):
		fmt.Println("YESS")
	default:
		fmt.Println("default")
	}
}
