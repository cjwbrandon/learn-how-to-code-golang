package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Women"},
		"no_dr":      {"Being evil", "Sunsets"},
	}

	x["abc"] = []string{"a", "b", "c"}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Println(k, v)
	}
}
