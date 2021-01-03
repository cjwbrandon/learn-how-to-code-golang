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

// INTERFACE
// Anyone that has a TYPE that has the METHOD speak is also HUMAN
// Note: A value can be of more than 1 type
type human interface {
	speak()
}

// Note: Empty interface -> Every value can be used
// type example interface {}
// func (a ...interface{}) // params can be of any type

func bar(h human) {
	fmt.Println("I called human")

	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar, (person)", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar, (secretAgent)", h.(secretAgent).first)
	}
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Mr",
		last:  "No",
	}

	fmt.Println(sa1)
	fmt.Println(p1)

	// Call method
	sa1.speak()
	// p1.speak()

	// Polymorphism -> Value can be of Type Human or secretAgent according to the function
	bar(sa1)
	bar(p1) // Error if person does not contain speak() method
}
