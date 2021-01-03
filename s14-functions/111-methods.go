package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Methods
// func (r receiver) {IDENTIFIER} (p params) returnType { ... }
// func is attached to receiver as a method
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)

	// Call method
	sa1.speak()
}
