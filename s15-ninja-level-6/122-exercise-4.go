package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi, my name is %v %v and I am %d years old.\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Bob",
		last:  "Jones",
		age:   34,
	}
	p1.speak()
}
