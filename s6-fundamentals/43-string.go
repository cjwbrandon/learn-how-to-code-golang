package main

import "fmt"

func main() {
	// Ways to declare a string " and `
	// Note: Strings are immutable -> you can reassign but not change the og string
	s1 := "Hello Playground"
	// Raw String literal
	s2 := `"Hello" 
	Playground`
	fmt.Println(s1, s2)
	fmt.Printf("%T\t%T\n", s1, s2)

	// Sequence of bytes -> slice of bytes
	// ASCII Scheme to represent letters of alphabets
	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// UTF-8 Code point -> rune
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println("")
	for i, v := range s1 {
		fmt.Printf("At index %d, we have %#x\n", i, v)
	}
}
