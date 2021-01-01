package main

import "fmt"

var y = 42

func main() {
	// Print
	fmt.Print(y)
	// Print with a new line
	fmt.Println(y)
	// Print with formating
	fmt.Printf("%v\n", y)
	fmt.Printf("Type: %T\n", y)  // type
	fmt.Printf("Bytes: %b\n", y) // Bytes
	fmt.Printf("Hexa-decimal: %x\n", y)
	fmt.Printf("Hexa-decimal with 0x: %#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// SPrint -> String Print
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// Fprint -> File Printing, printing to a file
}
