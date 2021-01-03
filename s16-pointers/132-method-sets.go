package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	// area2() float64 // Does not work -> Pointer receiver
}

// Non-pointer receiver
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Pointer receiver
func (c *circle) area2() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{
		radius: 4,
	}

	// Works with pointer and non-pointer values
	info(c1)
	info(&c1)
}
