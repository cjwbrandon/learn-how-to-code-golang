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

	persons := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(persons)

	for k, v := range persons {
		fmt.Println(k, v)
	}
}
