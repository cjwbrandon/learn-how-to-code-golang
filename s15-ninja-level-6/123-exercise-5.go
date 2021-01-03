package main

import (
	"fmt"
	"math"
)

type square struct {
	breadth float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return math.Pow(s.breadth, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		breadth: 4,
	}
	c1 := circle{
		radius: 4,
	}

	info(s1)
	info(c1)
}
