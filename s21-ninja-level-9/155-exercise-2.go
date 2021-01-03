package main

import "fmt"

type person struct {
	first string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("I am %v and I am %d years old.\n", p.first, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Bob",
		age:   23,
	}

	// saySomething(p1) // Error
	saySomething(&p1)
}
