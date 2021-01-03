package main

import "fmt"

func main() {
	bar(foo, "hkemlmlaos awmodrpladd")
	bar(moo, "hkemlmlaos awmodrpladd")
}

func foo(s string) {
	fmt.Println(s)
}

func moo(s string) {
	fmt.Println("Created from moo:", s)
}

func bar(f func(s string), s string) {
	var s2 string
	for i, l := range s {
		if i%2 == 0 {
			s2 += string(l)
		}
	}
	f(s2)
}
