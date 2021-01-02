package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)  // Add up using base 2 to get 72
	fmt.Printf("%#x\n", n) // Add up using base 16 to get 72
}
