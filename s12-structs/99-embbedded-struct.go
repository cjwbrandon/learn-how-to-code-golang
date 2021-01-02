package main

import "fmt"

type person struct {
	first string
	last  string
}

// Embedding in another type - sth like inheritance
type secretAgent struct {
	person
	first string // collision
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		first: "Show collision",
		ltk:   true,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(sa1, p2)
	fmt.Println(sa1.first, sa1.last, sa1.ltk) // Reference secretAgent's first property; else, inner type is promoted (i.e. last)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.ltk)
	fmt.Println(p2.first, p2.last)
}
