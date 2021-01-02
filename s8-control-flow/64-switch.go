package main

import "fmt"

func main() {
	// Switch with no expression == switch true {}
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print as well")
	case (3 == 3):
		fmt.Println("This should print")
		fallthrough
	case (4 == 4):
		fmt.Println("This only prints if you specify fallthrough")
	default:
		fmt.Println("this is default")
	}

	switch x := 42; x {
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
		fallthrough // to default
	default:
		fmt.Println("This is default")
	}

	y := "abc"
	switch y {
	// Multiple cases with same block
	case "a", "ab", "abc":
		fmt.Println("hahahha")
	default:
		fmt.Println("hehehe")
	}
}
