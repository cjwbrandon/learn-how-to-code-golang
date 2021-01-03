package main

import "fmt"

func main() {
	foo()
	bar("Bob")
	s1 := woo("Bob")
	fmt.Println(s1)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
}

// func(r receiver) identifier(parameters) (return(s)) { ... }
// func {IDENTIFIER}({PARAMS}) {RETURNTYPE} { ... }
func foo() {
	fmt.Println("Hello from foo")
}

// everything is Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hellow from bar,", s)
}

// Note: Parameters are function scoped
func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ", says hello")
	b := false
	return a, b
}
