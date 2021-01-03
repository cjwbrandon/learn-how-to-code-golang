package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	// *p = person{
	// 	first: "Changed",
	// 	last:  "Changed",
	// 	age:   123,
	// }
	(*p).first = "Changed"
}

func main() {
	p1 := person{
		first: "beforeChange",
		last:  "beforeChange",
		age:   1,
	}
	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}
