package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Women"},
		"no_dr":      {"Being evil", "Sunsets"},
	}

	for k, v := range x {
		fmt.Println(k, v)
	}
}
