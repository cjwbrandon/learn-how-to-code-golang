package main

import "fmt"

type person struct {
	first         string
	last          string
	favICFlavours []string
}

func main() {
	p1 := person{
		first:         "Bob",
		last:          "by",
		favICFlavours: []string{"choc", "vanilla"},
	}

	p2 := person{
		first:         "Kat",
		last:          "ty",
		favICFlavours: []string{"vanilla"},
	}

	fmt.Println(p1, p2)
}
